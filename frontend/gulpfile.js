var gulp = require('gulp')
    concat = require('gulp-concat'),
    uglify = require('gulp-uglify')
    rev = require('gulp-rev');

var build = '../server/frontend';

gulp.task('javascripts', function() {
  return gulp.src('javascripts/**/*.js')
    .pipe(concat('app.js'))
    .pipe(gulp.dest(build));
});

gulp.task('rev', ['javascripts'], function() {
  return gulp.src([build + '**/*.js'])
    .pipe(rev())
    .pipe(gulp.dest(build))
    .pipe(rev.manifest())
    .pipe(gulp.dest(build))
});

gulp.task('watch', ['javascripts'], function() {
  gulp.watch('javascripts/**/*.js', ['javascripts']);
});

gulp.task('default', ['javascripts']);
