
## Swagger Init
<details>
<summary> Install and Update Swagger data</summary>

```markdown
1. go install github.com/swaggo/swag/cmd/swag@v1.16.3
2. go get github.com/swaggo/gin-swagger
3. go get github.com/swaggo/swag
4. go get github.com/swaggo/files
5. swag init -g cmd/main.go
6. swag init -g cmd/main.go --parseDependency --parseInternal
```
</details>

This project is licensed under the Apache License 2.0. See the LICENSE file for details.