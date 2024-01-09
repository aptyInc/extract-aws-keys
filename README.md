# GitHub action to get bundle keys.

## usage

```cmd
      - name: Get split key
        uses: aptyInc/action-secret-exporter-go@master
        id: keys
        env:
          SECRETS: ${{ toJson(secrets) }}
          BRANCH_NAME: <name of the branch>
```
