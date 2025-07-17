angular.module('angular-go-blog').config(function($routeProvider) {
  $routeProvider
    .when('/', {
      templateUrl: 'app/views/home.html',
      controller: 'HomeController'
    })
    .when('/posts/:slug', {
      templateUrl: 'app/views/post.html',
      controller: 'PostController'
    })
    .otherwise({
      redirectTo: '/'
    });
});