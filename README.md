# whalefin, a Display Manager built with Wails

A minimal display manager for GNU/Systemd/Linux.

The goal is to create a very extensible, browser-based display manager.

```shell
make 
sudo make install
systemctl disable display-manager || true
systemctl enable whalefin
sudo reboot
```

That should be all you need if you are transitioning away from another Display
Manager. If you run into issues using whalefin, please visit 
[Troubleshooting](#Troubleshooting) below.

# Screenshot

The current version of whalefin looks like this:

![](img/screenshot.png)

# Development

The whalefin display manager can be tested by running as a normal Wails 
application (with `make run .`) or by running it within a Xephyr window 
(requires Xephyr installed) by running `make embed`.


# Troubleshooting

## Logging in Fails
### Problem
When attempting to log in, authentication fails despite using the correct
username and password.
### Solution
whalefin, like most other Display Managers, makes use of PAM, the Pluggable
Authentication Module system. This requires the appropriate rules be in place.
The following snippet creates a `display_manager` module by copying the login
one. It must be run as `root`:
`install -Dm00644 /etc/pam.d/login /etc/pam.d/display_manager`


## Audio Issues
### Problem
I'm using PulseAudio, and I had sound before using `startx`. Now, using 
whalefin, PulseAudio doesn't recognize any of my audio inputs or outputs.
 `pavucontrol` shows dummy sinks.

### Solution
whalefin uses `bash --login` to run your `.xinitrc`, and therefore imposes
stricter configuration requirements than `startx`. Create an `audio` group if
it does not already exist, and add your user to it. Log out and back in to 
reload the group list.


## Notifications No Longer Work
### Problem
DBUS is not initialized properly in the `.xinitrc` you're using. This is due
to your `XDG` variables (which `startx` handles for you) not being set properly.
### Solution
Be sure your XDG variables are set, such as
`export XDG_RUNTIME_DIR=/run/user/$(id -u)` towards the top of your `.xinitrc`.
Failing to do so breaks dbus and as a result, `systemctl` in user mode. Be sure
your `.xinitrc` is compliant with the recommendations of your distro, and
properly sources config files from `/etc/X11/xinit/xinitrc.d` (if applicable).

# Credits

The idea behind whalefin was loosely inspired by [fin](https://github.com/FyshOS/fin).

whalefin would not have been possible without the knowledge and expertise
of [Gulshan Singh](https://www.gulshansingh.com/posts/how-to-write-a-display-manager/).
Wails is © 2019-Present Lea Anthony.
