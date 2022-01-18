# op
"op" is a small tool that will allow you to open language or framework documentation in your browser from your terminal.

### Overview
It is not life changing, but it will help you save a few seconds here and there. Instead of leaving your terminal  or 
ide you can simply type `op <language|framework>` for a specific set of documentation, or you can type `op mdn <term>` 
to open up MDN docs with your query completed.

[Demo](https://www.loom.com/share/052024ea728645349a342c5e6607f5e7)

### Installation
This part is a work in progress that will hopefully become simpler over time.

### From source
If you write or use GoLang on your machine you can quite simply clone this repository

`git clone https://github.com/damiensedgwick/op`

Change in to the project directory and run `go install` - You *should* now be able to open a terminal anywhere you like
and run `op <language|framework>` or `op mdn <term>`

### Download executable binary from this repo
[Latest Releases](https://github.com/damiensedgwick/op/releases/latest)

The below notes need to be updated!! you will still need to give permissions to your executable but the naming should be okay now.

Once you have downloaded the appropiate executable, navigate to where you download the executable and do the following 
rename it, give it the correct permissions and move it to your bin which should also be in your path.

`mv <file_name> op`

`sudo chmod 755 op`

`mv op <bin|scripts>` - wherever you place your executables, it will need to be in your path to access it from anywhere.

You should now be able to run `op` from any terminal window on your system!

#### !Important for windows
You will have to make sure you do the windows equivalent of the above commands to get it working on your system correctly.

### Usage

In your terminal window, simply type either of the following.

`op <language|framework>` or `op mdn <query>`

#### NB.
When searching MDN, only the first word is currently used in the search query, this will update soon!

### Contributions
If there is a documentation link missing, and you would like to add it, feel free to open up a PR!
