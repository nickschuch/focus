# Focus

A tool to help focus by turning off Slack notifications for a set period of time.

This tool will:

* Pause Slack notifications for a given set of minutes.
* Set Slack status so people can see when the focus session will be complete so they can contact you if required.
* Let you know when the session is complete via system notifications.

## Example

```bash
$ export SLACK_TOKEN=xxxyyyzzz

$ focus 20
Pausing your notifications
Updating your status
Time to focus. Will send a notification when session is complete.
Session complete!
```
