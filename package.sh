#!/usr/bin/env bash

set -xe

go install fyne.io/fyne/v2/cmd/fyne@latest

echo "请选择类型";
# 选定项目
select env in "local" "macOS" "linux" "window" "android" "ios"
do
    case $env in
        "local")
            fyne install -icon myapp.png
            break
            ;;
        "macOS")
            fyne package -os darwin -icon myapp.png
            break
            ;;
        "linux")
            fyne package -os linux -icon myapp.png
            break
            ;;
        "window")
            fyne package -os windows -icon myapp.png
            break
            ;;
        "android")
            # 对于Android版本，您必须安装Android SDK和NDK，并设置适当的环境，以便在命令行上找到工具（如adb）
            # 要在手机或模拟器上安装android应用程序，只需调用：
            # adb install myapp.apk
            fyne package -os android -appID com.example.myapp -icon mobileIcon.png
            break
            ;;
        "ios")
            # 要在设备上安装iOS，请打开Xcode并在“窗口”菜单中选择“设备和模拟器”菜单项。然后找到你的手机，
            # 将myapp.app图标拖到你的应用列表中。要在模拟器上安装，可以使用以下命令行工具：
            # xcrun simctl install booted myapp.app
            fyne package -os ios -appID com.example.myapp -icon mobileIcon.png
            break
            ;;
        *)
            echo "输入错误，请重新输入"
    esac
done