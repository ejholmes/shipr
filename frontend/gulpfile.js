var gulp = require('gulp')
    concat = require('gulp-concat'),
    uglify = require('gulp-uglify');

gulp.task('javascripts', function() {
  return gulp.src('javascripts/**/*.js')
    .pipe(concat('app.js'))
    .pipe(gulp.dest('../server/frontend'));
});

gulp.task('watch', ['javascripts'], function() {
  gulp.watch('javascripts/**/*.js', ['javascripts']);
});

gulp.task('default', ['javascripts']);
