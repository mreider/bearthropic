X-callback-url Scheme documentation
Bear implements the x-callback-url protocol, which allow iOS and Mac developers to expose and document API methods they make available to other apps and return useful data.

Bear URL Scheme actions look like this:

bear://x-callback-url/[action]?[action parameters]&[x-callback parameters]

with x-success and x-error as available x-callback parameters.

Actions
/open-note
Open a note identified by its title or id and return its content.

parameters

id optional note unique identifier.
title optional note title.
header optional an header inside the note.
exclude_trashed optional if yes exclude trashed notes.
new_window optional if yes open the note in an external window (MacOS only).
float optional if yes makes the external window float on top (MacOS only).
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
open_note optional if no do not display the new note in Bear’s main or external window.
selected optional if yes return the note currently selected in Bear (token required)
pin optional if yes pin the note to the top of the list.
edit optional if yes place the cursor inside the note editor.
search optional opens the in-note find&replace panel with the specified text
x-success

note note text.
identifier note unique identifier.
title note title.
tags note tags array
is_trashed yes if the note is trashed.
modificationDate note modification date in ISO 8601 format.
creationDate note creation date in ISO 8601 format.
example

bear://x-callback-url/open-note?id=7E4B681B bear://x-callback-url/open-note?id=7E4B681B&header=Secondary%20Ttitle

Create and try /open-note actions in seconds with our URL builder online

/create
Create a new note and return its unique identifier. Empty notes are not allowed.

parameters

title optional note title.
text optional note body.
clipboard optional if yes use the text currently available in the clipboard
tags optional a comma separated list of tags.
file optional base64 representation of a file.
filename optional file name with extension. Both file and filename are required to successfully add a file.
open_note optional if no do not display the new note in Bear’s main or external window.
new_window optional if yes open the note in an external window (MacOS only).
float optional if yes make the external window float on top (MacOS only).
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
pin optional if yes pin the note to the top of the list.
edit optional if yes place the cursor inside the note editor.
timestamp optional if yes prepend the current date and time to the text
type optional if html the provided text parameter is converted from html to markdown
url optional if type is html this parameter is used to resolve relative image links
x-success

identifier note unique identifier.
title note title.
example

bear://x-callback-url/create?title=My%20Note%20Title&text=First%20line&tags=home,home%2Fgroceries

notes

The base64 file parameter have to be encoded when passed as an url parameter.

Create and try /create actions in seconds with our URL builder online

/add-text
append or prepend text to a note identified by its title or id. Encrypted notes can’t be accessed with this call.

parameters

id optional note unique identifier.
title optional title of the note.
selected optional if yes use the note currently selected in Bear (token required)
text optional text to add.
clipboard optional if yes use the text currently available in the clipboard
header optional if specified add the text to the corresponding header inside the note.
mode optional the allowed values are prepend, append, replace_all and replace (keep the note’s title untouched).
new_line optional if yes and mode is append force the text to appear on a new line inside the note
tags optional a comma separated list of tags.
exclude_trashed optional if yes exclude trashed notes.
open_note optional if no do not display the new note in Bear’s main or external window.
new_window optional if yes open the note in an external window (MacOS only).
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
edit optional if yes place the cursor inside the note editor.
timestamp optional if yes prepend the current date and time to the text
x-success

note note text.
title note title.
example

bear://x-callback-url/add-text?text=new%20line&id=4EDAF0D1&mode=append

Create and try /add-text actions in seconds with our URL builder online

/add-file
append or prepend a file to a note identified by its title or id. This call can’t be performed if the app is a locked state. Encrypted notes can’t be accessed with this call.

parameters

id optional note unique identifier.
title optional note title.
selected optional if yes use the note currently selected in Bear (token required)
file required base64 representation of a file.
header optional if specified add the file to the corresponding header inside the note.
filename required file name with extension. Both file and filename are required to successfully add a file.
mode optional the allowed values are prepend, append, replace_all and replace (keep the note’s title untouched).
open_note optional if no do not display the new note in Bear’s main or external window.
new_window optional if yes open the note in an external window (MacOS only).
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
edit optional if yes place the cursor inside the note editor.
x-success

note note text
example

bear://x-callback-url/add-file?filename=test.gif&id=4EDAF0D1-2EFF-4190-BC1D-67D9BAE49BA9-28433-000187BAA3D182EF&mode=append&file=R0lGODlhAQABAIAAAP%2F%2F%2F%2F%2F%2F%2FyH5BAEKAAEALAAAAAABAAEAAAICTAEAOw%3D%3D

notes

The base64 file parameter have to be encoded when passed as an url parameter.

Create and try /add-file actions in seconds with our URL builder online

/tags
Return all the tags currently displayed in Bear’s sidebar.

parameters

token required application token.
x-success

tags json array representing tags. [{ name }, ...]
example

bear://x-callback-url/tags?token=123456-123456-123456

Create and try /tags actions in seconds with our URL builder online

/open-tag
Show all the notes which have a selected tag in bear.

parameters

name required tag name or a list of tags divided by comma
token optional application token.
x-success

notes json array representing the tag’s notes. [{ title, identifier, modificationDate, creationDate, pin }, ...]
Encrypted notes will be excluded from the notes array. If more than one tag is passed with the name parameter this action returns all the notes matching one of the tags passed.

If token is not provided nothing is returned.

example

bear://x-callback-url/open-tag?name=work bear://x-callback-url/open-tag?name=todo%2Fwork

Create and try /open-tag actions in seconds with our URL builder online

/rename-tag
Rename an existing tag. This call can’t be performed if the app is a locked state. If the tag contains any locked note this call will not be performed.

parameters

name required tag name.
new_name required new tag name.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
example

bear://x-callback-url/rename-tag?name=todo&new_name=done

Create and try /rename-tag actions in seconds with our URL builder online

/delete-tag
Delete an existing tag. This call can’t be performed if the app is a locked state. If the tag contains any locked note this call will not be performed.

parameters

name required tag name.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
example

bear://x-callback-url/delete-tag?name=todo

Create and try /delete-tag actions in seconds with our URL builder online

/trash
Move a note to bear trash and select the Trash sidebar item. This call can’t be performed if the app is a locked state. Encrypted notes can’t be used with this call.

parameters

id optional note unique identifier.
search optional string to search.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
example

bear://x-callback-url/trash?id=7E4B681B bear://x-callback-url/trash?search=old

notes

The search term is ignored if an id is provided.

Create and try /trash actions in seconds with our URL builder online

/archive
Move a note to bear archive and select the Archive sidebar item. This call can’t be performed if the app is a locked state. Encrypted notes can’t be accessed with this call.

parameters

id optional note unique identifier.
search optional string to search.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
example

bear://x-callback-url/archive?id=7E4B681B bear://x-callback-url/archive?search=projects

notes

The search term is ignored if an id is provided.

Create and try /archive actions in seconds with our URL builder online

/untagged
Select the Untagged sidebar item.

parameters

search optional string to search.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
token optional application token.
x-success

notes json array representing the untagged notes. [{ title, identifier, [tag, ...], modificationDate, creationDate, pin }, ...]
Encrypted notes will be excluded from the notes array.

If token is not provided nothing is returned.

example

bear://x-callback-url/untagged?search=home

Create and try /untagged actions in seconds with our URL builder online

/todo
Select the Todo sidebar item.

parameters

search optional string to search.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
token optional application token.
x-success

notes json array representing the todo notes. [{ title, identifier, [tag, ...], modificationDate, creationDate, pin }, ...]
Encrypted notes will be excluded from the note array.

If token is not provided nothing is returned.

example

bear://x-callback-url/todo?search=home

Create and try /todo actions in seconds with our URL builder online

/today
Select the Today sidebar item.

parameters

search optional string to search.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
token optional application token.
x-success

notes json array representing the today notes. [{ title, identifier, [tag, ...], modificationDate, creationDate, pin }, ...]f
Encrypted notes will be excluded from the note array.

If token is not provided nothing is returned.

example

bear://x-callback-url/today?search=family

Create and try /today actions in seconds with our URL builder online

/locked
Select the Locked sidebar item.

parameters

search optional string to search.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
example

bear://x-callback-url/locked?search=data

Create and try /locked actions in seconds with our URL builder online

/search
Show search results in Bear for all notes or for a specific tag.

parameters

term optional string to search.
tag optional tag to search into.
show_window optional if no the call don’t force the opening of bear main window (MacOS only).
token optional application token.
x-success

notes json array representing the note results of the search. [{ title, identifier, [tag, ...], modificationDate, creationDate, pin }, ...]
Encrypted notes will be excluded from the note array.

If token is not provided nothing is returned.

example

bear://x-callback-url/search?term=nemo&tag=movies

Create and try /search actions in seconds with our URL builder online

/grab-url
Create a new note with the content of a web page.

parameters

url required url to grab.
tags optional a comma separated list of tags. If tags are specified in the Bear’s web content prefences this parameter is ignored.
pin optional if yes pin the note to the top of the list.
wait optional if no x-success is immediately called without identifier and title.
x-success

identifier note unique identifier.
title note title.
available values

yes no

example

bear://x-callback-url/grab-url?url=https://bear.app

Create and try /grab-url actions in seconds with our URL builder online

Token Generation
In order to extend their functionalties, some of the API calls allow an app generated token to be passed along with the other parameters. Please mind a Token generated on iOS is not valid for MacOS and vice-versa.

On MacOS, select Help → Advanced →API Token → Copy Token and will be available in your pasteboard.

On iOS go to the preferences → Advanced, locate the API Token section and tap the cell below to generate the token or copy it in your pasteboard.

Support
To discuss URL scheme improvements or reporting bugs please use our Support Form or Bear’s subreddit.

