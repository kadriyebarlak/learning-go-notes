package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func main() {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	_ = rdb.FlushDB(ctx).Err()

	// Add 1000 objects (products) to Redis
	for i := 0; i < 100; i++ {
		productKey := "product:" + strconv.Itoa(i)
		productValue := "Product Name " + strconv.Itoa(i)
		err := rdb.Set(ctx, productKey, productValue, 0).Err()
		if err != nil {
			fmt.Println("Failed to add to Redis:", err)
		}
	}

	// Create a Bloom Filter and add the products to it
	bloomKey := "bloom_filter:products"
	for i := 0; i < 100; i++ {
		productKey := "product:" + strconv.Itoa(i)
		_, err := rdb.Do(ctx, "BF.ADD", bloomKey, productKey).Result()
		if err != nil {
			fmt.Println("Failed to add to Bloom Filter:", err)
		}
	}

	// Check some keys in the Bloom Filter
	fmt.Println("Checking some products in the Bloom Filter:")
	for i := 0; i < 10; i++ {
		productKey := "product:" + strconv.Itoa(rand.Intn(1500)) // Random key between 0-1500
		exists, err := rdb.Do(ctx, "BF.EXISTS", bloomKey, productKey).Bool()
		if err != nil {
			fmt.Println("Bloom Filter query error:", err)
		} else if exists {
			fmt.Printf("%s exists in the Bloom Filter.\n", productKey)
		} else {
			fmt.Printf("%s does not exist in the Bloom Filter.\n", productKey)
		}
	}
}
