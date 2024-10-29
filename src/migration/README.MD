Install atlas:
curl -sSf https://atlasgo.sh | sh

Import models to ../migration/main.go

Generate:
atlas migrate diff --env local

Apply:
atlas migrate apply --env local