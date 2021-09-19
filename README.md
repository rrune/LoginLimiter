# Overview
Limit the duration a windows user can be logged in per day.  
  
The account you want to limit has a childrens microsoft account anyway? Try [Microsoft Family Safety](https://support.microsoft.com/en-us/account-billing/set-screen-time-limits-on-your-kids-devices-a593d725-fc4c-044c-284d-32eab0305ffd).  
Want something for Linux? Try [Timekpr-nExT](https://www.linuxuprising.com/2019/11/timekpr-next-is-linux-parental-control.html).  

# Features
- Limit the duration a windows user can be logged in per day  
- Individual duration for every day of the week
- Notifications

# Usage
Set the durations in minutes in `time.json`.  
  
Run the program on login by either using autostart (Win+R, "shell:startup") or task scheduler (remember to set the working directory).  
  
Optional: Run the [Tray Notifier](https://github.com/rrune/LoginLimiterNotifier).

# Building
Move to `src` and type `go build -ldflags "-H windowsgui"`

# Additional Info
- Adding time for today:  
  
  There is no in-program way to do this, so the only way to do is to stop it, manually edit `time.json` and restart the program (Program reads `time.json` once on startup, then just writes into it, so edits to `time.json` only go into effect after a restart, but will actually be overwritten after at most one minute without a restart)
  
- What happens if `time.json` get deleted?:
  
  It will just get created again on the next login with 0 minutes remaining
  
- Can this be circumvented?:
  
  In short: yes. Especially if the user is an administrator, there is no real way to prevent them from killing the process. But even then, knowing how to do it requires at least a little but of knowledge. In reality, the process will be buried somewhere in the task manager without an icon and with the name of the exe as its name. So if the exe is renamed to something weird like vifdnd.exe (very important file do not delete), it's likely a normal user won't immediately find and associate it with the Limiter. Pair it with the Notifier that sits in the Tray with a proper name and icon as a decoy and most normal users won't know the right process to kill. If the user is no administrator, there are ways to actually prevent the user from killing it.

- Pull Requests: 
  
  If you feel like you can improve this, or just want to fix my code (or the formatting of this Readme), feel free. I probably won't work on it further but will get notified about pull requests and will most likely approve them if they add value
