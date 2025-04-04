package main

const templateItemType = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/itemtypes.txt
package item

const (
{{- range $key, $value := . }}
    Type{{ replace $value.ItemType " " "" }} = "{{ $value.Code }}"
{{- end }}
)

var ItemTypes = map[string]Type{
{{- range $key, $value := . }}
    Type{{ replace $value.ItemType " " "" }}: {ID: {{ $key }}, Name: "{{ $value.ItemType }}", Code: "{{ $value.Code }}", Throwable: {{ if eq $value.Throwable "1" }}true{{ else }}false{{ end }}, Beltable: {{ if eq $value.Beltable "1" }}true{{ else }}false{{ end }}, BodyLocs: []LocationType{{ if $value.BodyLoc1 }}{ {{- $loc1 := $value.BodyLoc1 }}{{- $loc1 = replace $loc1 "tors" "Torso" }}{{- $loc1 = replace $loc1 "larm" "LeftArm" }}{{- $loc1 = replace $loc1 "rarm" "RightArm" }}{{- $loc1 = replace $loc1 "lrin" "LeftRing" }}{{- $loc1 = replace $loc1 "rrin" "RightRing" }}{{- $loc1 = replace $loc1 "glov" "Gloves" }}{{- $loc1 = replace $loc1 "feet" "Feet" }}{{- $loc1 = replace $loc1 "neck" "Neck" }}{{- $loc1 = replace $loc1 "head" "Head" }}{{- $loc1 = replace $loc1 "belt" "Belt" }}Loc{{ $loc1 }}{{ if and $value.BodyLoc2 (ne $value.BodyLoc2 $value.BodyLoc1) }}{{- $loc2 := $value.BodyLoc2 }}{{- $loc2 = replace $loc2 "tors" "Torso" }}{{- $loc2 = replace $loc2 "larm" "LeftArm" }}{{- $loc2 = replace $loc2 "rarm" "RightArm" }}{{- $loc2 = replace $loc2 "lrin" "LeftRing" }}{{- $loc2 = replace $loc2 "rrin" "RightRing" }}{{- $loc2 = replace $loc2 "glov" "Gloves" }}{{- $loc2 = replace $loc2 "feet" "Feet" }}{{- $loc2 = replace $loc2 "neck" "Neck" }}{{- $loc2 = replace $loc2 "head" "Head" }}{{- $loc2 = replace $loc2 "belt" "Belt" }}, Loc{{ $loc2 }}{{ end }}}{{ else }}{}{{ end }}},
{{- end }}
}`

const templateLevels = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/levels.txt
package area

var Areas = map[ID]Area{
{{- range $key, $value := . }}
    {{ $key }}: {Name: "{{ $value.LevelName }}", ID: {{ $key }}},
{{- end }}
}`

const templateSkillDesc = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/skilldesc.txt
package skill

var Desc = map[ID]Description{
{{- range $key, $value := . }}
    {{ $key }}: {Page: {{ $value.SkillPage }}, Row: {{ $value.SkillRow }}, Column: {{ $value.SkillColumn }}, ListRow: {{ $value.ListRow }}, IconCel: {{ $value.IconCel }}},
{{- end }}
}`

const templateSkills = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/skills.txt
package skill

var Skills = map[ID]Skill{
{{- range $key, $value := . }}
    {{ index $value "*Id" }}: {Name: "{{ $value.skill }}", ID: {{ $key }}, LeftSkill: {{ if eq $value.leftskill "1" }}true{{ else }}false{{ end }}, RightSkill: {{ if eq $value.rightskill "1" }}true{{ else }}false{{ end }}},
{{- end }}
}`

const templateWeapons = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/weapons.txt, cmd/txttocode/txt/armor.txt, cmd/txttocode/txt/misc.txt
package item

var Desc = map[int]Description{
{{- range $key, $value := . }}
    {{ $value.ID }}: {Name: "{{ $value.name }}", ID: {{ $value.ID }}, Code: "{{ $value.code }}", NormalCode: "{{ $value.normcode }}", UberCode: "{{ $value.ubercode }}", UltraCode: "{{ $value.ultracode }}", InventoryWidth: {{ $value.invwidth }}, InventoryHeight: {{ $value.invheight }}, MinDefense: {{ $value.minac }}, MaxDefense: {{ $value.maxac }}, MinDamage: {{ $value.mindam }}, MaxDamage: {{ $value.maxdam }}, TwoHandMinDamage: {{ index $value "2handmindam" }}, TwoHandMaxDamage: {{ index $value "2handmaxdam" }}, MinMissileDamage: {{ $value.minmisdam }}, MaxMissileDamage: {{ $value.maxmisdam }}, Speed: {{ $value.speed }}, StrengthBonus: {{ $value.StrBonus }}, DexterityBonus: {{ $value.DexBonus }}, RequiredStrength: {{ $value.reqstr }}, RequiredDexterity: {{ $value.reqdex }}, Durability: {{ $value.durability }}, RequiredLevel: {{ $value.levelreq }}, MaxSockets: {{ $value.gemsockets }}, Type: "{{ $value.type }}"},
{{- end }}`

const templateArmorAndMisc = `
{{- range $key, $value := . }}
    {{ $value.ID }}: {Name: "{{ $value.name }}", ID: {{ $value.ID }}, Code: "{{ $value.code }}", NormalCode: "{{ $value.normcode }}", UberCode: "{{ $value.ubercode }}", UltraCode: "{{ $value.ultracode }}", InventoryWidth: {{ $value.invwidth }}, InventoryHeight: {{ $value.invheight }}, MinDefense: {{ $value.minac }}, MaxDefense: {{ $value.maxac }}, MinDamage: {{ $value.mindam }}, MaxDamage: {{ $value.maxdam }}, TwoHandMinDamage: {{ index $value "2handmindam" }}, TwoHandMaxDamage: {{ index $value "2handmaxdam" }}, MinMissileDamage: {{ $value.minmisdam }}, MaxMissileDamage: {{ $value.maxmisdam }}, Speed: {{ $value.speed }}, StrengthBonus: {{ $value.StrBonus }}, DexterityBonus: {{ $value.DexBonus }}, RequiredStrength: {{ $value.reqstr }}, RequiredDexterity: {{ $value.reqdex }}, Durability: {{ $value.durability }}, RequiredLevel: {{ $value.levelreq }}, MaxSockets: {{ $value.gemsockets }}, Type: "{{ $value.type }}"},
{{- end }}
`

const templateObjects = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/objects.txt
package object

var Desc = map[int]Description{
{{- range $key, $value := . }}
	{{ index $value "*ID" }}: {Name: "{{ $value.Name }}", ID: {{ index $value "*ID" }}, SizeX: {{ $value.SizeX }}, SizeY: {{ $value.SizeY }}, Left: {{ $value.Left }}, Top: {{ $value.Top }}, Width: {{ $value.Width }}, Height: {{ $value.Height }}, Yoffset: {{ $value.Yoffset }}, Xoffset: {{ $value.Xoffset }}, HasCollision: {{ if eq $value.HasCollision0 "1" }}true{{ else }}false{{ end }}},
{{- end }}
}`

const templateEntrances = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/lvlwarp.txt
package entrance

var Desc = map[int]Description{
{{- range $key, $value := . }}
    {{ $value.UniqueId }}: {ID: {{ $value.UniqueId }}, Name: "{{ $value.Name }}", SelectX: {{ $value.SelectX }}, SelectY: {{ $value.SelectY }}, SelectDX: {{ $value.SelectDX }}, SelectDY: {{ $value.SelectDY }}, OffsetX: {{ $value.OffsetX }}, OffsetY: {{ $value.OffsetY }}, Direction: "{{ $value.Direction }}"},
{{- end }}
}`
const templateRarePrefixes = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/rareprefix.txt
package item

type RarePrefix struct {
    ID              int
    Name            string
    ItemTypes       []string
    ExcludeTypes    []string
}

var RarePrefixDesc = map[int]RarePrefix{
{{- range $index, $value := . }}
    {{ add $index 156 }}: {ID: {{ add $index 156 }}, Name: "{{ $value.name }}", ItemTypes: []string{ {{- $types := slice }}{{- range $i := iterate 7 }}{{- with index $value (printf "itype%d" (add $i 1)) }}{{- if ne . "" }}{{- $types = append $types . }}{{- end }}{{- end }}{{- end }}{{- range $i, $type := $types }}{{if $i}}, {{end}}"{{ $type }}"{{- end -}} }, ExcludeTypes: []string{ {{- $types := slice }}{{- range $i := iterate 4 }}{{- with index $value (printf "etype%d" (add $i 1)) }}{{- if ne . "" }}{{- $types = append $types . }}{{- end }}{{- end }}{{- end }}{{- range $i, $type := $types }}{{if $i}}, {{end}}"{{ $type }}"{{- end -}} }},
{{- end }}
}`
const templateRareSuffixes = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/raresuffix.txt
package item

type RareSuffix struct {
    ID              int
    Name            string
    ItemTypes       []string
}

var RareSuffixDesc = map[int]RareSuffix{
{{- range $index, $value := . }}
   {{ add $index 1 }}: {ID: {{ add $index 1 }}, Name: "{{ $value.name }}", ItemTypes: []string{ {{- $types := slice }}{{- range $i := iterate 7 }}{{- with index $value (printf "itype%d" (add $i 1)) }}{{- if ne . "" }}{{- $types = append $types . }}{{- end }}{{- end }}{{- end }}{{- range $i, $type := $types }}{{if $i}}, {{end}}"{{ $type }}"{{- end -}} }},
{{- end }}
}`
const templateMagicPrefixes = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/magicprefix.txt
package item

type MagicPrefix struct {
    ID              int
    Name            string
    Rare            bool
    Level           int
    LevelReq        int
    Group           int
    Mods            []Mod
    TransformColor  string
    ItemTypes       []string
    ExcludeTypes    []string
}

var MagicPrefixDesc = map[int]MagicPrefix{
{{- range $index, $value := . }}
{{- if ne $value.Name "" }}
    {{ add $index 747 }}: {ID: {{ add $index 747 }}, Name: "{{ $value.Name }}", Rare: {{ if eq $value.rare "1" }}true{{ else }}false{{ end }}, Level: {{ if eq $value.level "" }}0{{ else }}{{ $value.level }}{{ end }}, LevelReq: {{ if eq $value.levelreq "" }}0{{ else }}{{ $value.levelreq }}{{ end }}, Group: {{ if eq $value.group "" }}0{{ else }}{{ $value.group }}{{ end }}, Mods: []Mod{ {{- $mods := slice }}{{- if ne $value.mod1code "" }}{{- $mods = append $mods (printf "{Code: \"%s\", Param: \"%s\", Min: %s, Max: %s}" $value.mod1code $value.mod1param (default $value.mod1min "0") (default $value.mod1max "0")) }}{{- end }}{{- if ne $value.mod2code "" }}{{- $mods = append $mods (printf "{Code: \"%s\", Param: \"%s\", Min: %s, Max: %s}" $value.mod2code $value.mod2param (default $value.mod2min "0") (default $value.mod2max "0")) }}{{- end }}{{- if ne $value.mod3code "" }}{{- $mods = append $mods (printf "{Code: \"%s\", Param: \"%s\", Min: %s, Max: %s}" $value.mod3code $value.mod3param (default $value.mod3min "0") (default $value.mod3max "0")) }}{{- end }}{{- range $i, $mod := $mods }}{{if $i}}, {{end}}{{ $mod }}{{- end -}} }, TransformColor: "{{ $value.transformcolor }}", ItemTypes: []string{ {{- $types := slice }}{{- range $i := iterate 7 }}{{- with index $value (printf "itype%d" (add $i 1)) }}{{- if ne . "" }}{{- $types = append $types . }}{{- end }}{{- end }}{{- end }}{{- range $i, $type := $types }}{{if $i}}, {{end}}"{{ $type }}"{{- end -}} }, ExcludeTypes: []string{ {{- $types := slice }}{{- range $i := iterate 5 }}{{- with index $value (printf "etype%d" (add $i 1)) }}{{- if ne . "" }}{{- $types = append $types . }}{{- end }}{{- end }}{{- end }}{{- range $i, $type := $types }}{{if $i}}, {{end}}"{{ $type }}"{{- end -}} }},
{{- end }}
{{- end }}
}`

const templateMagicSuffixes = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/magicsuffix.txt
package item

type MagicSuffix struct {
    ID              int
    Name            string
    Rare            bool
    Level           int
    LevelReq        int
    Group           int
    Mods            []Mod
    TransformColor  string
    ItemTypes       []string
    ExcludeTypes    []string
}

type Mod struct {
    Code   string
    Param  string
    Min    int
    Max    int
}

var MagicSuffixDesc = map[int]MagicSuffix{
{{- range $index, $value := . }}
{{- if ne $value.Name "" }}
    {{ $index }}: {ID: {{ $index }}, Name: "{{ $value.Name }}", Rare: {{ if eq $value.rare "1" }}true{{ else }}false{{ end }}, Level: {{ if eq $value.level "" }}0{{ else }}{{ $value.level }}{{ end }}, LevelReq: {{ if eq $value.levelreq "" }}0{{ else }}{{ $value.levelreq }}{{ end }}, Group: {{ if eq $value.group "" }}0{{ else }}{{ $value.group }}{{ end }}, Mods: []Mod{ {{- $mods := slice }}{{- if ne $value.mod1code "" }}{{- $mods = append $mods (printf "{Code: \"%s\", Param: \"%s\", Min: %s, Max: %s}" $value.mod1code $value.mod1param (default $value.mod1min "0") (default $value.mod1max "0")) }}{{- end }}{{- if ne $value.mod2code "" }}{{- $mods = append $mods (printf "{Code: \"%s\", Param: \"%s\", Min: %s, Max: %s}" $value.mod2code $value.mod2param (default $value.mod2min "0") (default $value.mod2max "0")) }}{{- end }}{{- if ne $value.mod3code "" }}{{- $mods = append $mods (printf "{Code: \"%s\", Param: \"%s\", Min: %s, Max: %s}" $value.mod3code $value.mod3param (default $value.mod3min "0") (default $value.mod3max "0")) }}{{- end }}{{- range $i, $mod := $mods }}{{if $i}}, {{end}}{{ $mod }}{{- end -}} }, TransformColor: "{{ $value.transformcolor }}", ItemTypes: []string{ {{- $types := slice }}{{- range $i := iterate 7 }}{{- with index $value (printf "itype%d" (add $i 1)) }}{{- if ne . "" }}{{- $types = append $types . }}{{- end }}{{- end }}{{- end }}{{- range $i, $type := $types }}{{if $i}}, {{end}}"{{ $type }}"{{- end -}} }, ExcludeTypes: []string{ {{- $types := slice }}{{- range $i := iterate 5 }}{{- with index $value (printf "etype%d" (add $i 1)) }}{{- if ne . "" }}{{- $types = append $types . }}{{- end }}{{- end }}{{- end }}{{- range $i, $type := $types }}{{if $i}}, {{end}}"{{ $type }}"{{- end -}} }},
{{- end }}
{{- end }}
}`

const templateUniqueItems = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/uniqueitems.txt
package item

type UniqueName string

const (
    None UniqueName = ""
{{- range $key, $value := . }}
    {{- if and (ne $value.index "") (ne $value.enabled "0") (ne (index $value "lvl req") "") }}
    {{- if eq $value.index "Rainbow Facet" }}
        {{- if contains $value.prop4 "death-skill" }}
    RainbowFacet{{ replace $value.par4 " " "" }}Death UniqueName = "Rainbow Facet ({{ $value.par4 }} Death)"
        {{- else if contains $value.prop4 "levelup-skill" }}
    RainbowFacet{{ replace $value.par4 " " "" }}Level UniqueName = "Rainbow Facet ({{ $value.par4 }} Level)"
        {{- end }}
    {{- else }}
    {{ replace (replace (replace $value.index "'" "") " " "") "-" "" }} UniqueName = "{{ $value.index }}"
    {{- end }}
    {{- end }}
{{- end }}
)

type UniqueItemInfo struct {
    Name     string
    Code     string
    LevelReq int
    Rarity   int
    ID       int
}

var UniqueItems = map[UniqueName]UniqueItemInfo{
{{- range $key, $value := . }}
    {{- if and (ne $value.index "") (ne $value.enabled "0") (ne (index $value "lvl req") "") }}
    {{- if eq $value.index "Rainbow Facet" }}
        {{- if contains $value.prop4 "death-skill" }}
    RainbowFacet{{ replace $value.par4 " " "" }}Death: {Name: "Rainbow Facet", Code: "{{ $value.code }}", LevelReq: {{ index $value "lvl req" }}, Rarity: {{ $value.rarity }}, ID: {{ index $value "*ID" }}},
        {{- else if contains $value.prop4 "levelup-skill" }}
    RainbowFacet{{ replace $value.par4 " " "" }}Level: {Name: "Rainbow Facet", Code: "{{ $value.code }}", LevelReq: {{ index $value "lvl req" }}, Rarity: {{ $value.rarity }}, ID: {{ index $value "*ID" }}},
        {{- end }}
    {{- else }}
    {{ replace (replace (replace $value.index "'" "") " " "") "-" "" }}: {Name: "{{ $value.index }}", Code: "{{ $value.code }}", LevelReq: {{ index $value "lvl req" }}, Rarity: {{ $value.rarity }}, ID: {{ index $value "*ID" }}},
    {{- end }}
    {{- end }}
{{- end }}
}`

const templateSetItems = `// Code generated by cmd/txttocode. DO NOT EDIT.
// source: cmd/txttocode/txt/setitems.txt
package item

type SetName string
type SetItemName string

const (
    SetItemNone SetItemName = ""
{{- range $key, $value := . }}
    {{- if ne $value.index "" }}
    {{ replace (replace (replace $value.index "'" "") " " "") "-" "" }} SetItemName = "{{ $value.index }}"
    {{- end }}
{{- end }}
)

type SetItemInfo struct {
    Name      string
    SetName   string
    Code      string
    LevelReq  int
    Rarity    int
    ID        int
}

var SetItems = map[SetItemName]SetItemInfo{
{{- range $key, $value := . }}
    {{- if ne $value.index "" }}
    {{ replace (replace (replace $value.index "'" "") " " "") "-" "" }}: {Name: "{{ $value.index }}", SetName: "{{ $value.set }}", Code: "{{ $value.item }}", LevelReq: {{ index $value "lvl req" }}, Rarity: {{ $value.rarity }}, ID: {{ index $value "*ID" }}},
    {{- end }}
{{- end }}
}

var Sets = map[string][]SetItemName{
{{- $currentSet := "" }}
{{- $setItems := slice }}
{{- range $key, $value := . }}
    {{- if ne $value.index "" }}
        {{- if ne $currentSet $value.set }}
            {{- if ne $currentSet "" }}
    "{{ $currentSet }}": {
            {{- range $setItems }}
        {{ . }},
            {{- end }}
    },
            {{- end }}
            {{- $currentSet = $value.set }}
            {{- $setItems = slice }}
        {{- end }}
        {{- $setItems = append $setItems (printf "%s" (replace (replace (replace $value.index "'" "") " " "") "-" "")) }}
    {{- end }}
{{- end }}
{{- if ne $currentSet "" }}
    "{{ $currentSet }}": {
    {{- range $setItems }}
        {{ . }},
    {{- end }}
    },
{{- end }}
}`
