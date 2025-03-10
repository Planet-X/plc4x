<#--
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

	  https://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
-->
<#-- Prevent freemarker from escaping stuff -->
<#outputformat "undefined">
<#-- Declare the name and type of variables passed in to the template -->
<#-- @ftlvariable name="languageName" type="java.lang.String" -->
<#-- @ftlvariable name="protocolName" type="java.lang.String" -->
<#-- @ftlvariable name="outputFlavor" type="java.lang.String" -->
<#-- @ftlvariable name="helper" type="org.apache.plc4x.language.go.GoLanguageTemplateHelper" -->
<#-- @ftlvariable name="tracer" type="org.apache.plc4x.plugins.codegenerator.protocol.freemarker.Tracer" -->
<#-- @ftlvariable name="type" type="org.apache.plc4x.plugins.codegenerator.types.definitions.ComplexTypeDefinition" -->
${helper.fileName(protocolName, languageName, outputFlavor)?replace(".", "/")}/ParserHelper.go
/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package ${outputFlavor?replace("-","")}

import (
	"context"

	"github.com/apache/plc4x/plc4go/protocols/${helper.getSanitizedPackageName()}/${outputFlavor?replace("-","")}/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type ${helper.getSanitizedProtocolName()?cap_first}ParserHelper struct {
}

func (m ${helper.getSanitizedProtocolName()?cap_first}ParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (interface{}, error) {
	switch typeName {
<#list helper.getComplexTypeRootDefinitions() as typeDefinition>
	case "${typeDefinition.name}":
	<#if typeDefinition.parserArguments.isPresent()>
		<#list typeDefinition.parserArguments.orElseThrow() as parserArgument>
			<#if parserArgument.type.isSimpleTypeReference()>
		${parserArgument.name}, err := utils.StrTo${helper.getLanguageTypeNameForTypeReference(parserArgument.type)?cap_first}(arguments[${parserArgument?index}])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
			<#elseif parserArgument.type.isEnumTypeReference()>
        ${parserArgument.name}, _ := model.${helper.getLanguageTypeNameForTypeReference(parserArgument.type)}ByName(arguments[${parserArgument?index}])
			<#else>
		var ${parserArgument.name} model.${helper.getLanguageTypeNameForTypeReference(parserArgument.type)}
			</#if>
		</#list>
	</#if>
		return model.${typeDefinition.name}ParseWithBuffer(context.Background(), io<#if typeDefinition.parserArguments.isPresent()>, <#list typeDefinition.parserArguments.orElseThrow() as parserArgument>${parserArgument.name}<#sep>, </#list></#if>)
</#list>
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
</#outputformat>