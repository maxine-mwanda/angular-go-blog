angular.module('angular-go-blog.controllers').controller('PostController', 
  function($scope, $routeParams, BlogService) {
    $scope.loading = true;
    $scope.post = {};
    
    BlogService.getPost($routeParams.slug).then(function(post) {
      $scope.post = post;
      $scope.loading = false;
    }).catch(function(error) {
      console.error('Error fetching post:', error);
      $scope.loading = false;
    });
  }
);