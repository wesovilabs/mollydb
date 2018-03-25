'use strict';

var _codemirror = require('codemirror');

var _codemirror2 = _interopRequireDefault(_codemirror);

var _graphqlLanguageServiceInterface = require('graphql-language-service-interface');

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

/**
 * Registers a "hint" helper for CodeMirror.
 *
 * Using CodeMirror's "hint" addon: https://codemirror.net/demo/complete.html
 * Given an editor, this helper will take the token at the cursor and return a
 * list of suggested tokens.
 *
 * Options:
 *
 *   - schema: GraphQLSchema provides the hinter with positionally relevant info
 *
 * Additional Events:
 *
 *   - hasCompletion (codemirror, data, token) - signaled when the hinter has a
 *     new list of completion suggestions.
 *
 */
/**
 *  Copyright (c) 2015, Facebook, Inc.
 *  All rights reserved.
 *
 *  This source code is licensed under the BSD-style license found in the
 *  LICENSE file in the root directory of this source tree. An additional grant
 *  of patent rights can be found in the PATENTS file in the same directory.
 */

_codemirror2.default.registerHelper('hint', 'graphql', function (editor, options) {
  var schema = options.schema;
  if (!schema) {
    return;
  }

  var cur = editor.getCursor();
  var token = editor.getTokenAt(cur);
  var rawResults = (0, _graphqlLanguageServiceInterface.getAutocompleteSuggestions)(schema, editor.getValue(), cur, token);
  /**
   * GraphQL language service responds to the autocompletion request with
   * a different format:
   * type CompletionItem = {
   *   label: string,
   *   kind?: number,
   *   detail?: string,
   *   documentation?: string,
   *   // GraphQL Deprecation information
   *   isDeprecated?: ?string,
   *   deprecationReason?: ?string,
   * };
   *
   * Switch to codemirror-compliant format before returning results.
   */
  var tokenStart = token.type !== null && /"|\w/.test(token.string[0]) ? token.start : token.end;
  var results = {
    list: rawResults.map(function (item) {
      return {
        text: item.label,
        type: schema.getType(item.detail),
        description: item.documentation,
        isDeprecated: item.isDeprecated,
        deprecationReason: item.deprecationReason
      };
    }),
    from: { line: cur.line, column: tokenStart },
    to: { line: cur.line, column: token.end }
  };

  if (results && results.list && results.list.length > 0) {
    results.from = _codemirror2.default.Pos(results.from.line, results.from.column);
    results.to = _codemirror2.default.Pos(results.to.line, results.to.column);
    _codemirror2.default.signal(editor, 'hasCompletion', editor, results, token);
  }

  return results;
});