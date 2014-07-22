var gulp    = require('gulp'),
    concat  = require('gulp-concat'),
    uglify  = require('gulp-uglify'),
    sass    = require('gulp-sass'),
    html2js = require('gulp-ng-html2js'),
    merge   = require('merge-stream');

var build = '../server/frontend';

gulp.task('sass', function() {
  return gulp.src('stylesheets/**/*.scss')
    .pipe(sass())
    .pipe(gulp.dest(build))
});

gulp.task('stylesheets', ['sass']);

gulp.task('javascripts', function() {
  var templates = gulp.src('templates/**/*.html')
    .pipe(html2js({ moduleName: 'templates' }));

  var scripts = gulp.src('javascripts/**/*.js');

  return merge(scripts, templates)
    .pipe(concat('app.js'))
    .pipe(gulp.dest(build));
});

gulp.task('html', function() {
  return gulp.src('index.html')
    .pipe(gulp.dest(build));
});

gulp.task('watch', ['javascripts', 'stylesheets', 'html'], function() {
  gulp.watch('javascripts/**/*.js', ['javascripts']);
  gulp.watch('templates/**/*.html', ['javascripts']);
  gulp.watch('stylesheets/**/*.js', ['stylesheets']);
  gulp.watch('index.html', ['html']);
});

gulp.task('default', ['stylesheets', 'javascripts', 'html']);
