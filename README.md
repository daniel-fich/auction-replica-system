# About
This is a simulator of an auction system with clients to make bids and servers to receiver them. The goal of this is to try and create a service which will persist through server crashes.
# Run the launch script
```bash
sudo bash ./launch.sh
```
This will run the launch script which will clean the shared volume folder and startup docker compose. You can run it by hand by running the following.
```bash
rm -rf shared_volume/*
docker compose up --build
```
# Editing the compose file
If you want to change the amount of clients or servers you can change the replicas value in the compose file. But you should be beware of longer build times if increased.
```yaml
...
      replicas: 3
      ...
...
```
