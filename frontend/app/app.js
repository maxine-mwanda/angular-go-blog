angular.module('angular=go-blog', ['ngRoute'])
.config(['$locationProvider', function($locationProvider) {
    $locationProvider.hashPrefix('!');  // Important for compatibility
    $locationProvider.html5Mode({
        enabled: true,
        requireBase: true
    });
}]);angular.module('blogApp', ['ngRoute'])
.config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
    $locationProvider.hashPrefix('!');
    $locationProvider.html5Mode({
        enabled: true,
        requireBase: false  // Changed to false
    });
    
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
}]);