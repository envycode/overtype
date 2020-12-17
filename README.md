# OverType

Like [Typeracer](https://play.typeracer.com/) but for practicing multilanguage for example: hiragana -> romaji.

## Dev Dependencies

- Go 1.15
- Node
- Postgres
- Redis

## How to Run

### Backend
- run `make dep` to install dependencies
- run `make run`

### Web
- move to web directory `cd web`
- run `npm install`
- run `npm run dev`

## How to Deploy (for Contributors)
- Go to tab Actions on this repository
- Build your docker images with `workflow:build_and_push_registry`, please input your commit hash or branch name to build
- Deploy to overtype shared k8s cluster with `workflow:deploy_to_remote_server`, please input your commit hash or branch name that already build in previous stage

## How to Deploy (self hosted)
- You can directly use the Dockerfile to build your own image
- Please see `.env.example` to provide needed configuration
