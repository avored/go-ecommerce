let mix = require('laravel-mix');


/******** AVORED COPY IMAGES  **********/
mix.copyDirectory('resources/img', "public/img");


mix.js('resources/js/app.js', 'public/js/app.js').setPublicPath('public');
