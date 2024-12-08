package utils

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common/helpers"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// IFirebaseApp is an interface that has been extracted in order to support mocking of Firebase auth in tests
type IFirebaseApp interface {
	Auth(ctx context.Context) (*auth.Client, error)
	Firestore(ctx context.Context) (*firestore.Client, error)
	Messaging(ctx context.Context) (*messaging.Client, error)
}

// FirebaseClient is an implementation of the FirebaseClient interface
type FirebaseClient struct{}

// InitFirebase ensures that we have a working Firebase configuration
func (fc *FirebaseClient) InitFirebase() (IFirebaseApp, error) {
	appCreds, err := helpers.GetEnvVar(common.GoogleCloudProjectIDEnvVarName)
	if err != nil {
		return firebase.NewApp(
			context.Background(),
			nil,
			option.WithCredentialsFile(appCreds),
		)
	}
	return firebase.NewApp(
		context.Background(),
		nil,
	)
}

// GetFirestoreEnvironmentSuffix get the env suffix where the app is running
func GetFirestoreEnvironmentSuffix() string {
	return helpers.MustGetEnvVar("ROOT_COLLECTION_SUFFIX")
}

// ShortCodeSuffixCollection adds a shortcode's suffix to the collection name. This will aid in separating
// collections for different environments
func ShortCodeSuffixCollection(c string) string {
	return fmt.Sprintf("%v_sms_%v", c, GetFirestoreEnvironmentSuffix())
}

// USSDSuffixCollection adds a USSD's suffix to the collection name. This will aid in separating
// collections for different environments
func USSDSuffixCollection(c string) string {
	return fmt.Sprintf("%v_ussd_%v", c, GetFirestoreEnvironmentSuffix())
}

// DeleteCollection deletes firestore collection
func DeleteCollection(
	ctx context.Context,
	client *firestore.Client,
	ref *firestore.CollectionRef,
	batchSize int) error {
	for {
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}
