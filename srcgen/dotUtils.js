const changeCase = require('change-case')
const versionReg = /V.*\d/
const clientReg = /(Lor|Riot|Val|Lol|Tft)/
const optQueryParamsReg = /(^[a-z]+|[A-Z]+(?![a-z])|[A-Z][a-z]+)/

// flatMap: https://gist.github.com/samgiles/762ee337dff48623e729
// [B](f: (A) ⇒ [B]): [B]  ; Although the types in the arrays aren't strict (:
Array.prototype.flatMap = function (lambda) {
  return Array.prototype.concat.apply([], this.map(lambda))
}
Array.prototype.groupBy = function (lambda) {
  return Object.entries(
    this.reduce((agg, x) => {
      const k = lambda(x)
      ;(agg[k] = agg[k] || []).push(x)
      return agg
    }, {}),
  )
}
Array.prototype.sortBy = function (lambda) {
  return this.sort((a, b) => {
    const va = lambda(a)
    const vb = lambda(b)
    if (typeof va !== typeof vb) throw Error(`Mismatched sort types: ${typeof va}, ${typeof vb}.`)
    if (typeof va === 'number') return va - vb
    if (typeof va === 'string') return va.localeCompare(vb)
    throw Error(`Unknown sort type: ${typeof va}.`)
  })
}

function preamble(packageName, version) {
  return `\
// Automatically generated package.
package ${packageName}

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = ${version}`
}

function getEndpointGroupsByName(paths, name) {
  const endpointGroups = {}
  for (let path of Object.entries(paths)) {
    let api = path[0].split('/')[1]
    if (api !== name) continue
    let ep = path[1]['x-endpoint']
    endpointGroups[ep] = endpointGroups[ep] || []
    endpointGroups[ep].push(path)
  }
  return endpointGroups
}

function removeClientName(name) {
  return name.replace(clientReg, '')
}

function getNormalizedDTOStructName(name, version, endpoint) {
  let temp = name
  version = versionReg.exec(version)
  version = version === null ? null : version
  if (version[0].length !== 2) {
    version[0] = version[0].slice(version[0].length - 2, version[0].length)
  }
  temp = temp.replace('DTO', `${version[0]}DTO`)
  if (temp.includes(`${version[0]}${version[0]}DTO`)) {
    temp = temp.replace(`${version[0]}${version[0]}DTO`, `${version[0]}DTO`)
  }
  if (endpoint !== null && endpoint.includes('tournament') && endpoint.includes('stub')) {
    if (
      temp.startsWith('Tournament') ||
      temp.startsWith('LobbyEvent') ||
      temp.startsWith('Provider')
    ) {
      temp = 'Stub' + temp
    }
  }
  if (endpoint !== null) {
    if (endpoint.startsWith('league-exp')) {
      if (temp.startsWith('League') || temp.startsWith('Mini')) {
        temp = 'Exp' + temp
      }
    }
    if (endpoint.startsWith('val-ranked')) {
      if (temp.startsWith('Player')) {
        temp = 'Match' + temp
      }
    }
    if (endpoint.startsWith('lor-ranked')) {
      if (temp.startsWith('Player')) {
        temp = 'Leaderboard' + temp
      }
    }
    if (endpoint.startsWith('val-status')) {
      if (temp.startsWith('Content')) {
        temp = 'Status' + temp
      }
    }
  }
  if (temp.startsWith('ChampionInfoV')) {
    temp = temp.replace('ChampionInfoV', 'ChampionRotationV')
  }
  return temp
}

function getNormalizedFieldName(name) {
  let temp = name.replace('-', '')
  switch (temp) {
    case 'puuid':
      return 'PUUID'
    case 'xp':
      return 'XP'
    default:
      if (temp.endsWith('Id')) {
        temp = temp.slice(0, temp.length - 2) + 'ID'
      }
      if (temp.endsWith('Ids')) {
        temp = temp.slice(0, temp.length - 3) + 'IDs'
      }
      if (temp.includes('Ids')) {
        temp = temp.replace('Ids', 'IDs')
      }
      return capitalize(temp)
  }
}

function normalizeSchemaName(name) {
  if (!name.endsWith('DTO') && !name.endsWith('Dto')) {
    return name + 'DTO'
  }
  return name.replace('Dto', 'DTO')
}

function normalizeMethodName(method) {
  let temp = method
  if (temp.startsWith('Get')) {
    temp = temp.slice(3, temp.length)
  }
  if (temp.includes('League')) {
    temp = temp.replace('League', '')
  }
  if (temp.includes('Challenge') && temp !== 'Challenger') {
    temp = temp.replace('Challenge', '')
  }
  if (temp.endsWith('Id')) {
    temp = temp.slice(0, temp.length - 2) + 'ID'
  }
  if (temp.endsWith('Puuid')) {
    temp = temp.slice(0, temp.length - 5) + 'PUUID'
  }
  switch (temp) {
    case 'CurrentGameInfoBySummoner':
      return 'CurrentGameBySummonerID'
    case 'ChampionMasteryScoreByPUUID':
      return 'MasteryScoreByPUUID'
    case 'AllChampionMasteriesByPUUID':
      return 'AllMasteriesByPUUID'
    case 'ChampionMasteryByPUUID':
      return 'MasteryByPUUID'
    case 'TopChampionMasteriesByPUUID':
      return 'TopMasteriesByPUUID'
    case 'AllChampionMasteries':
      return 'AllMasteriesBySummonerID'
    case 'ChampionMastery':
      return 'MasteryBySummonerID'
    case 'TopChampionMasteries':
      return 'TopMasteriesBySummonerID'
    case 'ChampionMasteryScore':
      return 'ScoreBySummonerID'

    case 'PlayersByPUUID':
      return 'SummonerEntriesByPUUID'
    case 'PlayersBySummoner':
      return 'SummonerEntriesBySummonerID'
    case 'FeaturedGames':
      return 'Featured'

    case 'ShardData':
      return 'Shard'

    case 'TeamByID':
      return 'TeamByTeamID'

    case 'ChampionInfo':
      return 'Rotation'

    case 'BySummonerName':
      return 'ByName'

    case 'MatchIdsByPUUID':
      return 'ListByPUUID'
    case 'Matchlist':
      return 'ListByPUUID'

    case 'Configs':
      return 'ConfigByID'

    case 'PlayerData':
      return 'ByPUUID'

    case 'PlatformData':
      return 'Platform'

    case 'EntriesForSummoner':
      return 'SummonerEntries'

    case 'Challenger':
      return 'ChallengerByQueue'
    case 'Grandmaster':
      return 'GrandmasterByQueue'
    case 'Master':
      return 'MasterByQueue'

    case 'TournamentByID':
      return 'ByID'
    case 'TournamentByTeam':
      return 'ByTeamID'

    case 'Match':
      return 'ByID'
  }
  return temp
}

function capitalize(input) {
  return input[0].toUpperCase() + input.slice(1)
}

function normalizePropName(propName) {
  let out = propName
  // No leading digits.
  if (/^\d/.test(out)) {
    out = 'x' + out
  }
  if (out === 'Authorization') return out.toLowerCase()
  if (out === 'type') return out + '_'
  return out
}

function stringifyType(prop) {
  if (prop.anyOf) {
    prop = prop.anyOf[0]
  }

  let enumType = prop['x-enum']
  if (enumType && 'locale' !== enumType) {
    if (enumType === 'champion') return 'int'
    if (enumType === 'team') return 'int'
    return changeCase.pascalCase(enumType)
  }

  let refType = prop['$ref']
  if (refType) {
    return normalizeSchemaName(refType.slice(refType.indexOf('.') + 1))
  }

  switch (prop.type) {
    case 'boolean':
      return 'bool'
    case 'integer':
      return 'int32' === prop.format ? 'int32' : 'int64'
    case 'number':
      return 'float' === prop.format ? 'float32' : 'float64'
    case 'array':
      return '[]' + stringifyType(prop.items)
    case 'string':
      return 'string'
    case 'object':
      return `map[${stringifyType(prop['x-key'])}]` + stringifyType(prop.additionalProperties)
    default:
      return prop.type
  }
}

function formatJsonProperty(name) {
  return `\`json:"${name}"\``
}

function formatAddQueryParam(param) {
  const name = normalizePropName(param.name)
  const prop = param.schema
  if (prop.type === 'string') {
    return `if ${name} != \"\" {
    values.Set("${name}", fmt.Sprint(${name}))
  }`
  }
  if (prop.type === 'integer') {
    return `if ${name} != -1 {
    values.Set("${name}", fmt.Sprint(${name}))
  }`
  }
  throw `${prop.type} not supported`
}

function formatAddHeaderParam(param, returnValue, isPrimitive) {
  const name = normalizePropName(param.name)
  const prop = param.schema
  let value = `new(${returnValue})`
  if (isPrimitive) value = '*' + value
  if (prop.type === 'string') {
    return `if ${name} == \"\" {
    return ${value}, fmt.Errorf("'${name}' header is required")
  }
  request.Header.Set("${name}", fmt.Sprint(${name}))`
  }
  throw `${prop.type} not supported`
}

function formatRouteArgument(route, pathParams = []) {
  if (!pathParams.length) return `"${route}"`
  route = route.replace(/\{\S+?\}/g, '%v')
  const args = pathParams.map(({ name }) => name).join(', ')
  return `fmt.Sprintf("${route}", ${args})`
}

module.exports = {
  optQueryParamsReg,
  getEndpointGroupsByName,

  changeCase,
  preamble,
  capitalize,
  removeClientName,

  getNormalizedFieldName,
  getNormalizedDTOStructName,
  stringifyType,

  normalizeSchemaName,
  normalizePropName,
  normalizeMethodName,

  formatJsonProperty,
  formatAddQueryParam,
  formatAddHeaderParam,
  formatRouteArgument,
}
