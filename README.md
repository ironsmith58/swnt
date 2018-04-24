# swnt - Command line tools for Stars Without Number GMs

swnt provides a command line interface to the roll tables and generators found in Stars Without Number : Revised Edition (free version).
While the code is MIT licensed, all roll table content is the copyright of [Kevin Crawford, Sine Nominee Publishing](https://sinenominepublishing.com/).

## Usage

At present all commands are for generating content. Most useful is likely to be the Sector generator. Generating a new sector is done by issuing the following command:

    swnt new sector

You can optionally add the -l flag if you want cell text to be coloured according to Tech Level. When you choose to write a sector a new folder will be created and each system
will be provided with its own folder. This is designed to make it easy to correlate notes on a per system basis and keep things reasonably well organised.

All commands can be queried for their available options with the -h flag. For example:

    swnt new -h
    Generate content

    Usage:
      swntools new [command]

    Available Commands:
      adventure   Generate an Adventure
      conflict    Generate a Conflict
      encounter   Generate a quick encounter
      npc         Generate a NPC
      place       Generate a place
      poi         Generate a Point of Interest
      sector      Create the skeleton of a Sector
      world       Generate a secondary World for a Sector cell

    Flags:
      -h, --help   help for new

    Use "swntools new [command] --help" for more information about a command.

This is currently an alpha as I'm probably going to redesign some of the code for table layouts and a few other bits and pieces.