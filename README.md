# Astro Grid #
See https://wselwood.github.io/astro-grid/

Astro-grid is a simple display of the data contained in the minor planet center orbit files.

You can select which dimension to use on each axis using the select boxes at the top of the page.

The graph shows each dimension with the colour indicating the number of asteroids in that cell.

## Running locally ##
You will need a working go installation. You should be able to get this from your package manager or
the [golang website](http://golang.org/). Make sure your GOPATH is set.

You will also need a copy of the [minor planet center orbit file](http://minorplanetcenter.net/iau/MPCORB.html)

```
git clone https://github.com/wselwood/astro-grid.git
cd astro-grid
mkdir ./data
go build
./astro-grid -in $path_to_mpcorb.dat.gz -out ./data
```

Now open index.html in your browser.

## Project structure ##

`main.go` contains the main loop.

`dimensions.go` defines the dimensions. Each Dimension has an extractor which defines how
to get the data from a minor planet record.

`extractors.go` defines the extractors. This must define two things, how to find the cell for a given value
and how to find the base value for that cell. Tests are in `extractors_test.go`

`grid.go` contains the data structures that back the result grids while processing.

`index.html` contains the rendering code for the visualization. This uses D3.


## Contributions ##

All contributions are warmly welcomed. Please report bugs using the bug tracker the git hub project.
Please fork, branch and raise pull requests when submitting code changes. Any questions or
