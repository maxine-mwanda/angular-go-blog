angular.module('angular-go-blog.controllers')
.controller('HomeController', ['$scope', '$location', 'BlogService', 
function($scope, $location, BlogService) {
    $scope.posts = [];
    $scope.loading = true;
    
    BlogService.getPosts().then(function(posts) {
        $scope.posts = posts;
        $scope.loading = false;
    });
    
    $scope.navigateToPost = function(slug) {
        $location.path('/posts/' + slug);
    };
}]);