(function(module) {
try {
  module = angular.module('templates');
} catch (e) {
  module = angular.module('templates', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('content.html',
    '<div class="container" ui-view></div>\n' +
    '');
}]);
})();

// ansi_up.js
// version : 1.0.0
// author : Dru Nelson
// license : MIT
// http://github.com/drudru/ansi_up

(function (Date, undefined) {

    var ansi_up,
        VERSION = "1.0.0",

        // check for nodeJS
        hasModule = (typeof module !== 'undefined'),

        // Normal and then Bright
        ANSI_COLORS = [
          ["0,0,0", "187, 0, 0", "0, 187, 0", "187, 187, 0", "0, 0, 187", "187, 0, 187", "0, 187, 187", "255,255,255" ],
          ["85,85,85", "255, 85, 85", "0, 255, 0", "255, 255, 85", "85, 85, 255", "255, 85, 255", "85, 255, 255", "255,255,255" ]
        ];

    function Ansi_Up() {
      this.fg = this.bg = null;
      this.bright = 0;
    }

    Ansi_Up.prototype.escape_for_html = function (txt) {
      return txt.replace(/[&<>]/gm, function(str) {
        if (str == "&") return "&amp;";
        if (str == "<") return "&lt;";
        if (str == ">") return "&gt;";
      });
    };

    Ansi_Up.prototype.linkify = function (txt) {
      return txt.replace(/(https?:\/\/[^\s]+)/gm, function(str) {
        return "<a href=\"" + str + "\">" + str + "</a>";
      });
    };

    Ansi_Up.prototype.ansi_to_html = function (txt) {

      var data4 = txt.split(/\033\[/);

      var first = data4.shift(); // the first chunk is not the result of the split

      var self = this;
      var data5 = data4.map(function (chunk) {
        return self.process_chunk(chunk);
      });

      data5.unshift(first);

      var flattened_data = data5.reduce( function (a, b) {
        if (Array.isArray(b))
          return a.concat(b);

        a.push(b);
        return a;
      }, []);

      var escaped_data = flattened_data.join('');

      return escaped_data;
    };

    Ansi_Up.prototype.process_chunk = function (text) {

      // Do proper handling of sequences (aka - injest vi split(';') into state machine
      //match,codes,txt = text.match(/([\d;]+)m(.*)/m);
      var matches = text.match(/([\d;]+?)m([^]*)/m);

      if (!matches) return text;

      var orig_txt = matches[2];
      var nums = matches[1].split(';');

      var self = this;
      nums.map(function (num_str) {

        var num = parseInt(num_str);

        if (num === 0) {
          self.fg = self.bg = null;
          self.bright = 0;
        } else if (num === 1) {
          self.bright = 1;
        } else if ((num >= 30) && (num < 38)) {
          self.fg = "rgb(" + ANSI_COLORS[self.bright][(num % 10)] + ")";
        } else if (num === 39) {
          self.fg = null;
        } else if ((num >= 40) && (num < 48)) {
          self.bg = "rgb(" + ANSI_COLORS[0][(num % 10)] + ")";
        }
      });

      if ((self.fg === null) && (self.bg === null)) {
        return orig_txt;
      } else {
        var style = [];
        if (self.fg)
          style.push("color:" + self.fg);
        if (self.bg)
          style.push("background-color:" + self.bg);
        return ["<span style=\"" + style.join(';') + "\">", orig_txt, "</span>"];
      }
    };

    // Module exports
    ansi_up = {

      escape_for_html: function (txt) {
        var a2h = new Ansi_Up();
        return a2h.escape_for_html(txt);
      },

      linkify: function (txt) {
        var a2h = new Ansi_Up();
        return a2h.linkify(txt);
      },

      ansi_to_html: function (txt) {
        var a2h = new Ansi_Up();
        return a2h.ansi_to_html(txt);
      },

      ansi_to_html_obj: function () {
        return new Ansi_Up();
      }
    };

    // CommonJS module is defined
    if (hasModule) {
        module.exports = ansi_up;
    }
    /*global ender:false */
    if (typeof window !== 'undefined' && typeof ender === 'undefined') {
        window.ansi_up = ansi_up;
    }
    /*global define:false */
    if (typeof define === "function" && define.amd) {
        define("ansi_up", [], function () {
            return ansi_up;
        });
    }
})(Date);

(function(module) {
try {
  module = angular.module('templates');
} catch (e) {
  module = angular.module('templates', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('header.html',
    '\n' +
    '<div class="navbar navbar-default navbar-static-top" id="header">\n' +
    '  <div class="container">\n' +
    '    <div class="navbar-header">\n' +
    '      <a class="navbar-brand" href="" ui-sref="app.jobs.list">Shipr</a>\n' +
    '    </div>\n' +
    '    <div class="navbar-collapse collapse">\n' +
    '      <ul class="nav navbar-nav navbar-right">\n' +
    '        <li class="user">\n' +
    '        <a href="" ng-bind="user.username"></a>\n' +
    '        <img class="gravatar" ng-src="{{user.gravatar}}">\n' +
    '        </li>\n' +
    '      </ul>\n' +
    '    </div>\n' +
    '  </div>\n' +
    '</div>\n' +
    '');
}]);
})();

(function(angular) {
  'use strict';

  var module = angular.module('app', [
    'ng',
    'ngSanitize',
    'ui.router',
    'templates',
    'app.directives',
    'app.services',
    'app.filters',
    'app.controllers'
  ]);

  module.config(function($locationProvider, $stateProvider) {
    $locationProvider.html5Mode(true);

    $stateProvider
      .state('app', {
        'abstract': true,
        views: {
          header: { templateUrl: 'header.html' },
          content: { templateUrl: 'content.html' }
        }
      })

      .state('app.jobs', {
        'abstract': true,
        templateUrl: 'jobs.html'
      })

      .state('app.jobs.list', {
        url: '/',
        controller: 'JobsListCtrl',
        templateUrl: 'jobs/list.html',
        resolve: {
          jobs: function(Job) {
            return Job.all();
          }
        }
      })

      .state('app.jobs.detail', {
        url: '/deploys/:jobId',
        controller: 'JobsDetailCtrl',
        templateUrl: 'jobs/detail.html',
        resolve: {
          job: function($stateParams, Job) {
            return Job.find($stateParams.jobId);
          }
        }
      });
  });

  module.config(function($httpProvider) {
    $httpProvider.defaults.headers.common = {
      'Accept': 'application/vnd.shipr+json; version=1'
    }
  });

  module.run(function($rootScope, $log) {
    $rootScope.$on('$stateChangeError', function(event, toState, toParams, fromState, fromParams, error) {
      $log.error(error);
    });
  });

})(angular);

(function(module) {
try {
  module = angular.module('templates');
} catch (e) {
  module = angular.module('templates', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('jobs.html',
    '<div ui-view></div>\n' +
    '');
}]);
})();

(function(angular) {
  'use strict';

  var module = angular.module('app.controllers', [
    'ng'
  ]);

  module.controller('JobsListCtrl', function($scope, jobs) {
    $scope.jobs = jobs;
  });

  module.controller('JobsDetailCtrl', function($scope, $state, job) {
    $scope.job = job;

    $scope.restart = function() {
      $scope.job.restart().then(function(job) {
        $state.go('app.jobs.detail', { jobId: job.id });
      });
    };
  });

})(angular);

(function(module) {
try {
  module = angular.module('templates');
} catch (e) {
  module = angular.module('templates', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('jobs/detail.html',
    '<div class="job">\n' +
    '  <header class="job__header">\n' +
    '    <h1 class="job__status">\n' +
    '      <span ng-class="{ \'is-queued\': job.isQueued(), \'is-deploying\': job.isDeploying() }" ng-if="job.isQueued() || job.isDeploying()">\n' +
    '        <div class="spinner"></div>\n' +
    '        <span ng-if="job.isQueued()">Queued</span>\n' +
    '        <span ng-if="job.isDeploying()">Deploying</span>\n' +
    '      </span>\n' +
    '      <span class="is-done" ng-if="job.isDeployed()">Deployed</span>\n' +
    '      <span class="is-failed" ng-if="job.isFailed()">Failed</span>\n' +
    '    </h1>\n' +
    '    <p class="job__destination">\n' +
    '    <strong ng-bind="job.ref" title="{{ job.sha }}"></strong>\n' +
    '    &#8674;\n' +
    '    <strong ng-bind="job.environment"></strong>\n' +
    '    </p>\n' +
    '  </header>\n' +
    '  <div class="panel panel-primary" ng-if="job.config">\n' +
    '    <div class="panel-heading">\n' +
    '      <span ng-bind="job.repo.name"></span>\n' +
    '      <a class="restart btn btn-info btn-xs pull-right" href="" ng-click="restart()" ng-if="job.done">Restart</a>\n' +
    '    </div>\n' +
    '    <div class="panel-body">\n' +
    '      <div environment-variables="job.config"></div>\n' +
    '    </div>\n' +
    '  </div>\n' +
    '  <div class="job__not-started" ng-if="!job.isStarted()">\n' +
    '    <strong>Hang Tight!</strong>\n' +
    '    Your logs should be showing up shortly\n' +
    '  </div>\n' +
    '  <div class="job__log" id="log" ng-bind-html="job.output | ansi" ng-if="job.isStarted()" sticky="job.output"></div>\n' +
    '</div>\n' +
    '');
}]);
})();

(function(angular) {
  'use strict';

  var module = angular.module('app.directives', [
    'ng'
  ]);

  /**
   * A directive for building a css3 spinner.
   */
  module.directive('spinner', function() {
    return {
      restrict: 'C',
      link: function(scope, elem) {
        function addRect(i) {
          elem.append('<div class="rect' + i + '"></div> ');
        }

        _(4).times(addRect);
      }
    };
  });

  module.directive('sticky', function($document, $window) {
    var padding = 100;

    return {
      restrict: 'A',
      link: function(scope, elem, attrs) {
        var $doc = $window.$($document),
            $win = $window.$($window);

        scope.$watch(attrs.sticky, function() {
          var sticky = $doc.scrollTop() + $win.height() >= $doc.height() - padding;

          if (sticky) {
            $doc.scrollTop($doc.height());
          }
        });
      }
    }
  });

  /**
   * A directive that for showing environment variables.
   */
  module.directive('environmentVariables', function($compile) {
    return {
      restrict: 'A',
      scope: { environmentVariables: '=' },
      link: function(scope, elem) {
        _.each(scope.environmentVariables, function(value, key) {
          elem.append($compile('<span environment-variable var="' + key + '" value="' + value + '" />')(scope));
        });
      }
    };
  });

  module.directive('environmentVariable', function() {
    return {
      restrict: 'EA',
      scope: { var: '@', value: '@' },
      template: '<div class="environment-variable"><span class="var" ng-bind="var"></span>=<span class="value" ng-bind="value"></span></div>'
    };
  });

})(angular);

(function(module) {
try {
  module = angular.module('templates');
} catch (e) {
  module = angular.module('templates', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('jobs/list.html',
    '<h2 class="page-header">Deploys</h2>\n' +
    '<table class="table table-striped">\n' +
    '  <thead>\n' +
    '    <tr>\n' +
    '      <th>#</th>\n' +
    '      <th>Sha</th>\n' +
    '    </tr>\n' +
    '  </thead>\n' +
    '  <tbody>\n' +
    '  <tr ng-repeat="job in jobs | orderBy:\'id\':true">\n' +
    '    <td>\n' +
    '      <a href="" ng-bind="job.id" ui-sref="app.jobs.detail({ jobId: job.id })"></a>\n' +
    '    </td>\n' +
    '    <td ng-bind="job.sha"></td>\n' +
    '  </tr>\n' +
    '  </tbody>\n' +
    '</table>\n' +
    '');
}]);
})();

(function(angular) {
  'use strict';

  var module = angular.module('app.filters', [
    'ng'
  ]);

  module.filter('ansi', function($window, $sce) {
    var ansi_up = $window.ansi_up,
        ansi_to_html = ansi_up.ansi_to_html,
        escape_for_html = ansi_up.escape_for_html;

    return function(input) {
      return $sce.trustAsHtml(ansi_to_html(escape_for_html(input)));
    };
  });

})(angular);

(function(angular) {
  'use strict';

  var module = angular.module('app.services', [
    'ng',
    'ngResource'
  ]);

  module.factory('Job', function($resource) {
    var resource = $resource(
      '/jobs/:jobId',
      { jobId: '@id' },
      { restart: { method: 'POST', url: '/jobs/:jobId/restart' } }
    );

    function Job(attributes){
      this.setAttributes(attributes);

      this.restart = attributes.$restart;
    }

    /**
     * Get a single job.
     */
    Job.find = function(id) {
      return resource.get({ jobId: id }).$promise.then(function(job) {
        return new Job(job);
      });
    };

    /**
     * Get all jobs.
     */
    Job.all = function() {
      return resource.query().$promise;
    };

    _.extend(Job.prototype, {
      /**
       * Set the attributes on this model.
       *
       * @param {Object} attributes
       */
      setAttributes: function(attributes) {
        var job = this;

        _.each(attributes, function(value, key) {
          job[key] = value;
        });
      },

      /**
       * Append some log output.
       *
       * @param {String} output
       */
      appendOutput: function(output) {
        this.output += output;
      },

      /**
       * Whether or not the job has started to be worked on.
       *
       * @return {Boolean}
       */
      isStarted: function() {
        return this.status != 'pending';
      },

      /**
       * Whether or not the job is queueud.
       *
       * @return {Boolean}
       */
      isQueued: function() {
        return !this.isStarted();
      },

      /**
       * Whether or not the job is deploying.
       *
       * @return {Boolean}
       */
      isDeploying: function() {
        return !this.done && this.isStarted();
      },

      /**
       * Whether or not the job successfully deployed.
       *
       * @return {Boolean}
       */
      isDeployed: function() {
        return this.done && this.success;
      },

      /**
       * Whether or not the job failed to deploy.
       *
       * @return {Boolean}
       */
      isFailed: function() {
        return this.done && !this.success;
      }
    });

    return Job;
  });

})(angular);
