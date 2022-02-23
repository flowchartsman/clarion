tell application "System Events"
    if exists (window "[Extension Development Host] - build_themes.go — mktheme" of process "Code") then
        set clarionWindow to (window "[Extension Development Host] - build_themes.go — mktheme" of process "Code")
        perform action "AXRaise" of clarionWindow
        set position of clarionWindow to {0, 0}
    else
        error "Theme debug window not found"
    end if
end tell
activate application "Code"
tell application "System Events"
    tell process "Code"
        set the size of front window to {1300,900}
    end tell
end tell
tell application "System Events"
   keystroke "p" using {command down, shift down}
   delay 1
   keystroke "t"
   keystroke "e"
   keystroke "r"
   keystroke "m"
   keystroke "i"
   keystroke "n"
   keystroke "a"
   keystroke "l"
   keystroke "c"
   keystroke "r"
   keystroke "e"
   key code 36
   delay 3
   keystroke "."
   keystroke "/"
   keystroke "b"
   keystroke "u"
   keystroke "i"
   keystroke "l"
   keystroke "d"
   keystroke "/"
   keystroke "c"
   keystroke "o"
   keystroke "l"
   keystroke "o"
   keystroke "r"
   keystroke "s"
   keystroke "."
   keystroke "s"
   keystroke "h"
   key code 36
end tell
--do shell script "screencapture -x -R0,23,900,900 \"foo.png\""
--tell application "System Events"
--    tell process "Code"
--        click at {100,100}
--    end tell
--end tell
--tell application "Finder" to get the bounds of clarionWindow
