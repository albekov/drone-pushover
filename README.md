# drone-pushover

Send push notifications to your device using [Pushover](https://pushover.net/) from [Drone](https://drone.io/) pipelines.

## Usage

```yaml
kind: pipeline
type: docker
name: default

steps:
  - name: notify-start
    image: albekov/drone-pushover
    settings:
      user:
        from_secret: pushover_user
      token:
        from_secret: pushover_token
      message: |
        Started build #{{ build.number }} of {{ repo.name }} (type: "{{ build.event }}")
      title: |
        Branch {{ build.branch }}
```

## Configuration

- `user` - Pushover user key
- `token` - Pushover application token
- `message` - Notification message
- `title` - Notification title (optional)
- `device` - Pushover device name (optional)
