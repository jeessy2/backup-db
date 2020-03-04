# backup databases
  Only support postgres right now

  ```
  # change config
  vi config.yml
  # build docker and run
  docker build . -t dbback-up
  docker run -d dbback-up
  ```