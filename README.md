# GitHub action to get keys

# usage 

```yaml
jobs:
  getKeys:
    name: GetKeys
    runs-on: ubuntu-latest
    steps:
      - name: Get Keys
        uses: aptyInc/extract-aws-keys@v2
        id: KEYS
        env:
          SECRETS: ${{ toJson(secrets) }}
          AWS_REGION: "us-east-1"
          ENVIRONMENT: "development"
```

```sh
git tag -a v2 -m "Release version 2" 
git push origin v2  
```