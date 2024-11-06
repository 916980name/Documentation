### Oh no! Something went wrong
UI broken, need to `Ctrl + Alt + F2` enter command line.
1. Check network
    
    - curl 'something'

        Network connection failed.
    - Check hardware

        The network card light is off.
        > Unplug network cable and plug in again.
    - curl 'something' again, make sure net work

1. (Optional) Network config wrong

    - Go to `/etc/netplan`, check the yaml is correct. 
        > `netplan` command. You can not choose specific config yaml.
    - `nmcli` command, choose network profile
        > `nmcli c show`

        > `nmcli c up [profile name]`

1. The opeation solved my problem
    ```
    sudo apt-get update && sudo apt-get upgrade
    sudo dpkg --configure -a
    sudo apt-get clean && sudo apt-get autoremove
    sudo reboot
    ```
    > https://dev.to/justplegend/error-after-upgrade-on-ubuntu-2404-oh-no-something-went-wrong-1840

### The mouse scroll wheel is not working well
### The keyboard key mapping is incorrect
Tab, direction key.  
> These two problems recovered automaticly after I am using about 10 minutes. Maybe the driver installed automaticly.

Actually it is like the `shift` key been pressing, input something with `shift` key could work.  
Use left & right `shift`.

Resolved:
> https://askubuntu.com/questions/1300264/text-is-deleted-when-double-click-the-input-box-in-ubuntu-20-04

remove package: `ibus-sunpinyin`  
install package: `ibus-libpinyin`