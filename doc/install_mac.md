# Installation on MacOS

At this stage, the MacOS FLEcli binary is not signed. It will fail to load on recent MacOS version.

This procedure will guide you for adding FLEcli to the exception list.

When the executable has been unpacked in a directory and that you try to execute it in the console (`./FLEcli`), you will get the following error:

![](pictures/Error_1.png?raw=true)

Make sure that you choose "cancel" to close the dialog.
Having triggered that error will allow you to enable the exception.

Enter the "system preference" and choose the "Security Settings".

![](pictures/System_preference.png?raw=true)

On the "general" tab, you will see the problem that just occurred ( _"FLEcli" was blocked from use because it is not from an identified developer_ ). Click on the "Allow Anyway" button.

![](pictures/security_setting.png?raw=true)

When trying again to execute the application, the error message will be slightly different:

![](pictures/Error_2.png?raw=true)

But this time, choose "open". The exception is now registered.