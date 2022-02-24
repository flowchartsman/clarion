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
   -- kill all old terminals
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
   keystroke "k"
   keystroke "i"
   keystroke "l"
   keystroke "l"
   keystroke "a"
   keystroke "l"
   keystroke "l"
   key code 36
   delay 3
   -- Create new terminal for colors
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
   keystroke "s"
   keystroke "c"
   keystroke "r"
   keystroke "e"
   keystroke "e"
   keystroke "n"
   keystroke "s"
   keystroke "h"
   keystroke "o"
   keystroke "t"
   keystroke "_"
   keystroke "s"
   keystroke "c"
   keystroke "r"
   keystroke "i"
   keystroke "p"
   keystroke "t"
   keystroke "s"
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
   -- switch back to the editor
   delay 1
   key code 18 using control down
end tell
--do shell script "screencapture -x -R0,23,900,900 \"foo.png\""
--tell application "System Events"
--    tell process "Code"
--        click at {100,100}
--    end tell
--end tell
--tell application "Finder" to get the bounds of clarionWindow
