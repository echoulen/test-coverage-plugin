# test-coverage-plugin
A test coverage reported plugin with Drone CI

```
pipeline:
  test:
    image: node:9.3-alpine
    commands:
      - yarn test 2>&1 | tee output
      - yarn build
  report:
    image: echoulen/test-coverage-plugin:latest
    file: output
    repo: My/TestRepo
    host: 127.0.0.1
    port: 3200
    network_mode: host
```
