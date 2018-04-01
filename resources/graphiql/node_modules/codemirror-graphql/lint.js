'use strict';

var _codemirror = require('codemirror');

var _codemirror2 = _interopRequireDefault(_codemirror);

var _graphqlLanguageServiceInterface = require('graphql-language-service-interface');

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

/**
 *  Copyright (c) 2015, Facebook, Inc.
 *  All rights reserved.
 *
 *  This source code is licensed under the BSD-style license found in the
 *  LICENSE file in the root directory of this source tree. An additional grant
 *  of patent rights can be found in the PATENTS file in the same directory.
 */

var SEVERITY = ['error', 'warning', 'information', 'hint'];
var TYPE = {
  'GraphQL: Validation': 'validation',
  'GraphQL: Deprecation': 'deprecation',
  'GraphQL: Syntax': 'syntax'
};

/**
 * Registers a "lint" helper for CodeMirror.
 *
 * Using CodeMirror's "lint" addon: https://codemirror.net/demo/lint.html
 * Given the text within an editor, this helper will take that text and return
 * a list of linter issues, derived from GraphQL's parse and validate steps.
 * Also, this uses `graphql-language-service-parser` to power the diagnostics
 * service.
 *
 * Options:
 *
 *   - schema: GraphQLSchema provides the linter with positionally relevant info
 *
 */
_codemirror2.default.registerHelper('lint', 'graphql', function (text, options) {
  var schema = options.schema;
  var rawResults = (0, _graphqlLanguageServiceInterface.getDiagnostics)(text, schema);

  var results = rawResults.map(function (error) {
    return {
      message: error.message,
      severity: SEVERITY[error.severity - 1],
      type: TYPE[error.source],
      from: _codemirror2.default.Pos(error.range.start.line, error.range.start.character),
      to: _codemirror2.default.Pos(error.range.end.line, error.range.end.character)
    };
  });

  return results;
});