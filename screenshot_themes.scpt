tell application "System Events"
    set allWindows to name of window of process whose visible is true
end tell
return allWindows

set theWindowList to my subListsToOneList(myList) --  Flattening a list of lists
return theWindowList

on subListsToOneList(l)
    set newL to {}
    repeat with i in l
        set newL to newL & i
    end repeat
    return newL
end subListsToOneList
