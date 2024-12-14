gomark - CLI Bookmark Manager
gomark is a simple, fast, and lightweight command-line tool for managing bookmarks with the ability to store, search, and organize your bookmarks using plain text files. The tool supports adding bookmarks with custom shortcuts, listing them, and pruning outdated shortcuts. You can also open bookmarks via these shortcuts or search terms, and integrate it with fzf for easy searching.

Features
Add Bookmarks: Save bookmarks with custom labels or shortcuts.
Manage Shortcuts: Use shortcuts for faster access to frequently used links.
Prune Shortcuts: Clean up shortcuts that no longer correspond to saved bookmarks.
Search and Open Bookmarks: Easily search and open bookmarks using fzf or directly through shortcuts.
Fully POSIX Compatible: Works well with Unix-based systems, providing a simple and efficient workflow without needing a browser.
Installation
To install gomark, follow the instructions based on your system.

Using Go (Preferred)
Install Go (if you don't have it already):
Download and install Go from here.

Clone the repository:

bash
Copy code
git clone https://github.com/your-username/gomark.git
cd gomark
Build the binary:

bash
Copy code
go build -o gomark
Move the binary to your executable path:

bash
Copy code
mv gomark /usr/local/bin/gomark
Commands
Add a Bookmark
To add a bookmark, use the gomark add command. You can optionally assign a shortcut to the bookmark.

bash
Copy code
gomark add <link> -g <group> -s <shortcut>
-g <group>: Assigns the bookmark to a group (e.g., work, personal).
-s <shortcut>: Assigns a shortcut for easy reference.
Example:

bash
Copy code
gomark add https://meet.google.com/foo-bar-baz -g work -s standup
This adds the bookmark https://meet.google.com/foo-bar-baz to the work group and assigns the shortcut standup.

List Shortcuts
To list all the current shortcuts:

bash
Copy code
gomark shortcuts
Prune Shortcuts
To remove shortcuts that no longer have a corresponding bookmark:

bash
Copy code
gomark prune
This command scans the .shortcuts file and removes shortcuts that don't match any existing bookmarks.

Open a Bookmark by Shortcut or Search Term
To open a bookmark, you can either use its assigned shortcut or search term:

bash
Copy code
gomark open <shortcut or search-term>
If the term is a shortcut, it will open the corresponding URL.
If it's a search term, gomark will search through your saved bookmarks and present options for you to choose from using fzf (or simply open if only one match exists).
Example:

bash
Copy code
gomark open standup
This opens the bookmark associated with the standup shortcut.

File Structure
gomark uses simple plain-text files for organizing bookmarks:

Bookmarks: Stored in files under ~/.gomarks/<group>.txt (one file per group).
Shortcuts: Stored in a single .shortcuts file (~/.gomarks/.shortcuts).
Example:

javascript
Copy code
~/.gomarks/
  ├── work/
  │   └── project-a.txt
  └── .shortcuts
~/.gomarks/work/project-a.txt: Contains the list of bookmarks for the work group.
~/.gomarks/.shortcuts: Contains all the shortcuts in the format shortcut <url>.
Example Workflow
Add a bookmark:

bash
Copy code
gomark add https://google.com -g personal -s google
List your shortcuts:

bash
Copy code
gomark shortcuts
Output:

arduino
Copy code
google https://google.com
Open a bookmark via shortcut:

bash
Copy code
gomark open google
Prune invalid shortcuts:

bash
Copy code
gomark prune
Dependencies
Go: Required to build and run gomark.
fzf (optional): Used for interactive searching through bookmarks when multiple results are found.
You can install fzf using a package manager like apt, brew, or from the official GitHub page: fzf GitHub.

Contributing
Contributions are welcome! Feel free to submit issues, pull requests, or suggest improvements.

License
This project is licensed under the MIT License. See the LICENSE file for more details.

Contact
For questions or feedback, you can reach out via GitHub Issues.
