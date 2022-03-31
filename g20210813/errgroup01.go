package main

// 并发编程包之 errgroup
// https://mp.weixin.qq.com/s/NcrENqRyK9dYrOBBI0SGkA
import (
	"context"
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"
)

type (
	Result string
	Search func(ctx context.Context, query string) (Result, error)
)

func fakeSearch(kind string) Search {
	return func(_ context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}

func main() {
	Google := func(ctx context.Context, query string) ([]Result, error) {
		g, ctx := errgroup.WithContext(ctx)

		searches := []Search{Web, Image, Video}
		results := make([]Result, len(searches))

		for i, search := range searches {
			i, search := i, search
			g.Go(func() error {
				result, err := search(ctx, query)
				if err != nil {
					results[i] = result
				}
				return err
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}

		return results, nil
	}

	results, err := Google(context.Background(), "golang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)
