# DevOpsDays India - Bengaluru Chapter - Website #

## About ##

This project powers the [DevOpsDays India Website](https://devopsdaysindia.org/) (Bengaluru chapter). It uses the Project Zeppelin / GDG DevFest 2014 site template.

## About the site template ##

It is built on top of [Jekyll](http://jekyllrb.com/) - simple, blog-aware, static site generator. Template is brought by [GDG Lviv](http://lviv.gdg.org.ua/) team. Live demo here: http://gdg-x.github.io/zeppelin/. The automated version with Grunt: <https://github.com/gdg-x/zeppelin-grunt>

For a detailed walkthrough, watch the project presentation from [GDG[x] Townhall meeting](http://www.youtube.com/watch?v=xYmhheoLjcI). Slides available [here](https://docs.google.com/presentation/d/19aM7yNl_orDaCNND5LpCY3fShb6PyMltnzYfKvV8R_8/edit?usp=sharing)


## Local development ##

Check if you have [all requirements for local environment](http://jekyllrb.com/docs/installation/).

Recommended Ruby version: **3.1.2**

To install all development dependencies install [Bundler](http://bundler.io/).
```bash
gem install bundler
```

and run next command from root folder:
```bash
bundle install
```  

To start Jekyll run:
```bash
jekyll serve -w --baseurl /test
```

Site will be available at http://127.0.0.1:4000/test/ or http://localhost:4000/test/ (on Windows)

**NOTE:** in this mode all changes to html and data files will be automatically regenerated, but after changing ```_config.yml``` you have to restart server.


### Sass (Compass) support ###
**Note:** You need to install [Node.js](http://nodejs.org/download/)

To watch changes of `.sass` files and compile it to the `.css` on a fly change property `safe: true` to `safe: false` in `_config.yml`.
**Note: It works only on local machine, because GitHub runs Jekyll in `--save` [mode](https://help.github.com/articles/using-jekyll-with-pages/#configuration-overrides)**

Learn more about Sass development from this [documentation](https://github.com/gdg-x/zeppelin/wiki/Sass-development) by the GDG team.


### Resource optimizations (optional) ###

You can optimize images and minify css and javascript automatically (for now only on Windows).
But for Mac OS users available amazing tool - [imageoptim](https://imageoptim.com/). Thanks [@raphaelsavina](https://github.com/raphaelsavina) for link.
Optimize all images by running this script from `/automation/images/` folder:
```bash
all_image_optimization.bat -d -jtran -pout -pquant -optip -gsicle -svgo
```

To minify CSS and JS run `minify_js.bat` (for Windows) and `minify_js.sh` (for Linux and MacOS) from `/automation/minifying/` folder:
```bash
minify_js.bat
```

Learn more about available optimization options from this [documentation](https://github.com/gdg-x/zeppelin/wiki/Resources-optimizations) by the GDG team.


## Documentation ##

Quick-start guide is not enough? Checkout [full documentation](https://github.com/gdg-x/zeppelin/wiki) provided by the GDG team.


## Contributors ##

#### DevOpsDays India - Bengaluru - Website ####

See [list of contributors](https://github.com/emergingtechs/devopsdays/graphs/contributors)

#### Project Zeppelin / GDG DevFest 2014 Site Template ####

* Design and web development: [Oleh Zasadnyy](https://github.com/ozasadnyy)
* Idea: [Vitaliy Zasadnyy](https://github.com/zasadnyy)

See [list of contributors](https://github.com/gdg-x/zepplin/graphs/contributors)

Maintainers: [@tasomaniac](https://github.com/tasomaniac) and [@ozasadnyy](https://github.com/ozasadnyy).


## License ##

Project is published under the [MIT license](https://github.com/emergingtechs/devopsdays/blob/gh-pages/LICENSE.txt). Feel free to clone and modify repo as you want, but don't forget to add reference to authors :)
