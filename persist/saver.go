package persist

import (
	"context"

	"gopkg.in/olivere/elastic.v5"
)

//ItemSave ...
func ItemSave(item interface{}) error {
	//Must turn off sniff in docker

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return err
	}

	_, err = client.Index().
		Index("house_info").
		Type("rent").
		BodyJson(item).
		Do(context.Background())

	return err

}
