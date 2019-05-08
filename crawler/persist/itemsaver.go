package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"go-crawler/crawler/types"
	"github.com/kataras/iris/core/errors"
)

func ItemSaver(index string) (chan types.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan types.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client,index string, item types.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	//fmt.Printf("%+v", resp)
	return err
}

