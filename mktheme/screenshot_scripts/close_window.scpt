tell application "System Events"
    if exists (window "[Extension Development Host] - build_themes.go — mktheme" of process "Code") then
        set clarionWindow to (window "[Extension Development Host] - build_themes.go — mktheme" of process "Code")
        perform action "AXRaise" of clarionWindow
        set position of clarionWindow to {0, 0}
    else
        error "Theme debug window not found"
    end if
end tell
tell application "System Events"
   keystroke "q" using command down
end tell
