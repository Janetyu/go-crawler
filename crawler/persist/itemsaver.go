package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("item Saver: got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	//fmt.Printf("%+v", resp)
	return resp.Id, err
}

