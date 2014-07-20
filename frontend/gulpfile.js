var gulp = require('gulp')
    concat = require('gulp-concat'),
    uglify = require('gulp-uglify'),
    sass = require('gulp-sass');

var build = '../server/frontend';

gulp.task('sass', function() {
  return gulp.src('stylesheets/**/*.scss')
    .pipe(sass())
    .pipe(gulp.dest(build))
});

gulp.task('stylesheets', ['sass']);

gulp.task('javascripts', function() {
  return gulp.src('javascripts/**/*.js')
    .pipe(concat('app.js'))
    .pipe(gulp.dest(build));
});

gulp.task('watch', ['javascripts'], function() {
  gulp.watch('javascripts/**/*.js', ['javascripts']);
});

gulp.task('default', ['stylesheets', 'javascripts']);
