# Trading212 to Stockopedia

This allows you to export a CSV from trading212.com and import it directly into a Folio on stockopedia.com

It requires your Stockopedia username & password to map the securities into Google Tickers.

# Usage

### Container

```bash
docker run --rm --env-file stocko.env -v $(pwd):/tmp \
  gsdevme/trading212-to-stockopedia:latest convert \
  --file /tmp/from_2021-02-01_to_2021-11-01_MTYzNTgyNTYwMjE2Mg.csv
```

```bash
./trading212-to-stockopedia convert -h
Usage:
  Trading212 convert [flags]

Flags:
      --file string       The path to the CSV to convert
  -h, --help              help for convert
      --password string   Password for stockopedia
      --username string   Username for stockopedia
```