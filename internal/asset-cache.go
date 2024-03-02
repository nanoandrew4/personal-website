package internal

import (
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"personal-website/internal/logging"
	"strings"
)

var (
	hashForAsset = map[string]string{}
)

func init() {
	cssAssets, err := os.ReadDir("assets/css")
	if err != nil {
		log.Fatalln("cannot find assets, server should not start")
	}

	for _, entry := range cssAssets {
		if strings.Contains(entry.Name(), ".css") {
			cssAssetContent, readFileErr := os.ReadFile(fmt.Sprintf("assets/css/%s", entry.Name()))
			if readFileErr != nil {
				log.Fatalf("cannot read file %s for hash generation", entry.Name())
			}
			hash := fnv.New32a()
			_, err = hash.Write(cssAssetContent)
			if err != nil {
				log.Fatalf("error hashing file %s", entry.Name())
			}
			hashForAsset[fmt.Sprintf("/assets/css/%s", entry.Name())] = fmt.Sprintf("%x", hash.Sum32())
		}
	}
}

func GetPathWithQueryHash(assetPath string) string {
	fileHash, fileHashPresent := hashForAsset[assetPath]
	if !fileHashPresent {
		logging.GlobalLogger.Warn("Hash not available for asset", "asset", assetPath)
		return assetPath
	}
	return fmt.Sprintf("%s?%s", assetPath, fileHash)
}
