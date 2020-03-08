'use strict';

goog.module('grrUi.flow.flowResultsDirective');
goog.module.declareLegacyNamespace();



/**
 * Controller for FlowResultsDirective.
 *
 * @constructor
 * @param {!angular.Scope} $scope
 * @ngInject
 */
const FlowResultsController = function($scope) {
  /** @private {!angular.Scope} */
  this.scope_ = $scope;

  /** @type {?string} */
  this.flowResultsUrl;

  /** @type {?string} */
    this.flowExportedResultsUrl;

  /** @type {string} */
  this.exportBasePath;

  /** @type {?string} */
  this.outputPluginsMetadataUrl;

  /** @type {?string} */
    this.downloadFilesUrl;

    /** @type {object} */
    this.queryParams;

  /** @type {?string} */
    this.exportCommand;

    /** @type {string} */
    this.flowId;

    this.scope_.$watchGroup(
        ['flowId', 'controller.selectedArtifact', 'artifactNames'],
        this.onFlowIdOrBasePathChange_.bind(this));
};


/**
 * Handles directive's arguments changes.
 *
 * @param {Array<string>} newValues
 * @private
 */
FlowResultsController.prototype.onFlowIdOrBasePathChange_ = function(
    newValues, oldValues) {

   if (angular.isDefined(this.scope_.artifactNames) &&
        this.scope_.artifactNames.length > 0 &&
        this.selectedArtifact == null) {
        this.selectedArtifact = this.scope_.artifactNames[0];
    }

    if (newValues != oldValues) {
        if (newValues.every(angular.isDefined)) {
            var components = this.selectedArtifact.split("/");
            // Artifact has no source name.
            if (components.length == 1) {
                this.queryParams = {
                    path: '/clients/'+ this.scope_['clientId'] +
                        '/artifacts/' + this.selectedArtifact +
                        '/' + this.scope_.flowId + '.csv',
                    client_id: this.scope_['clientId'],
                };
            } else {
                // Artifact has a source name.
                this.queryParams = {
                    path: '/clients/' + this.scope_['clientId'] +
                        '/artifacts/' + components[0] +
                        '/' + this.scope_.flowId + "/" +
                        components[1] + '.csv',
                    client_id: this.scope_['clientId'],
                };
            }
        }
    }
};


/**
 * Directive for displaying results of a flow with a given URL.
 *
 * @return {!angular.Directive} Directive definition object.
 * @ngInject
 * @export
 */
exports.FlowResultsDirective = function() {
  return {
    scope: {
        artifactNames: '=',
        flowId: '=',
        clientId: '=',
    },
    restrict: 'E',
    templateUrl: '/static/angular-components/flow/flow-results.html',
    controller: FlowResultsController,
    controllerAs: 'controller'
  };
};


/**
 * Directive's name in Angular.
 *
 * @const
 * @export
 */
exports.FlowResultsDirective.directive_name = 'grrFlowResults';
