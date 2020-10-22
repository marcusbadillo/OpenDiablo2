package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// [https://d2mods.info/forum/viewtopic.php?t=41556, https://d2mods.info/forum/kb/viewarticle?a=246]

// SkillDetails has all of the SkillRecords
type SkillDetails map[int]*SkillRecord

// SkillRecord is a row from the skills.txt file. Here are two resources for more info on each field
type SkillRecord struct {
	Skill             string
	Charclass         string
	Skilldesc         string
	Prgcalc1          d2calculation.Calculation
	Prgcalc2          d2calculation.Calculation
	Prgcalc3          d2calculation.Calculation
	Srvmissile        string
	Srvmissilea       string
	Srvmissileb       string
	Srvmissilec       string
	Srvoverlay        string
	Aurastate         string
	Auratargetstate   string
	Auralencalc       d2calculation.Calculation
	Aurarangecalc     d2calculation.Calculation
	Aurastat1         string
	Aurastatcalc1     d2calculation.Calculation
	Aurastat2         string
	Aurastatcalc2     d2calculation.Calculation
	Aurastat3         string
	Aurastatcalc3     d2calculation.Calculation
	Aurastat4         string
	Aurastatcalc4     d2calculation.Calculation
	Aurastat5         string
	Aurastatcalc5     d2calculation.Calculation
	Aurastat6         string
	Aurastatcalc6     d2calculation.Calculation
	Auraevent1        string
	Auraevent2        string
	Auraevent3        string
	Auratgtevent      string
	Auratgteventfunc  string
	Passivestate      string
	Passiveitype      string
	Passivestat1      string
	Passivecalc1      d2calculation.Calculation
	Passivestat2      string
	Passivecalc2      d2calculation.Calculation
	Passivestat3      string
	Passivecalc3      d2calculation.Calculation
	Passivestat4      string
	Passivecalc4      d2calculation.Calculation
	Passivestat5      string
	Passivecalc5      d2calculation.Calculation
	Passiveevent      string
	Passiveeventfunc  string
	Summon            string
	Pettype           string
	Petmax            d2calculation.Calculation
	Summode           string
	Sumskill1         string
	Sumsk1calc        d2calculation.Calculation
	Sumskill2         string
	Sumsk2calc        d2calculation.Calculation
	Sumskill3         string
	Sumsk3calc        d2calculation.Calculation
	Sumskill4         string
	Sumsk4calc        d2calculation.Calculation
	Sumskill5         string
	Sumsk5calc        d2calculation.Calculation
	Sumoverlay        string
	Stsound           string
	Stsoundclass      string
	Dosound           string
	DosoundA          string
	DosoundB          string
	Tgtoverlay        string
	Tgtsound          string
	Prgoverlay        string
	Prgsound          string
	Castoverlay       string
	Cltoverlaya       string
	Cltoverlayb       string
	Cltmissile        string
	Cltmissilea       string
	Cltmissileb       string
	Cltmissilec       string
	Cltmissiled       string
	Cltcalc1          d2calculation.Calculation
	Cltcalc2          d2calculation.Calculation
	Cltcalc3          d2calculation.Calculation
	Range             string
	Itypea1           string
	Itypea2           string
	Itypea3           string
	Etypea1           string
	Etypea2           string
	Itypeb1           string
	Itypeb2           string
	Itypeb3           string
	Etypeb1           string
	Etypeb2           string
	Anim              d2enum.PlayerAnimationMode
	Seqtrans          string
	Monanim           string
	ItemCastSound     string
	ItemCastOverlay   string
	Skpoints          d2calculation.Calculation
	Reqskill1         string
	Reqskill2         string
	Reqskill3         string
	State1            string
	State2            string
	State3            string
	Perdelay          d2calculation.Calculation
	Calc1             d2calculation.Calculation
	Calc2             d2calculation.Calculation
	Calc3             d2calculation.Calculation
	Calc4             d2calculation.Calculation
	ToHitCalc         d2calculation.Calculation
	DmgSymPerCalc     d2calculation.Calculation
	EType             string
	EDmgSymPerCalc    d2calculation.Calculation
	ELenSymPerCalc    d2calculation.Calculation
	ID                int
	Srvstfunc         int
	Srvdofunc         int
	Srvprgfunc1       int
	Srvprgfunc2       int
	Srvprgfunc3       int
	Prgdam            int
	Aurafilter        int
	Auraeventfunc1    int
	Auraeventfunc2    int
	Auraeventfunc3    int
	Sumumod           int
	Cltstfunc         int
	Cltdofunc         int
	Cltprgfunc1       int
	Cltprgfunc2       int
	Cltprgfunc3       int
	Attackrank        int
	Weapsel           int
	Seqnum            int
	Seqinput          int
	LineOfSight       int
	SelectProc        int
	ItemEffect        int
	ItemCltEffect     int
	ItemTgtDo         int
	ItemTarget        int
	Reqlevel          int
	Maxlvl            int
	Reqstr            int
	Reqdex            int
	Reqint            int
	Reqvit            int
	Restrict          int
	Delay             int
	Checkfunc         int
	Startmana         int
	Minmana           int
	Manashift         int
	Mana              int
	Lvlmana           int
	Param1            int
	Param2            int
	Param3            int
	Param4            int
	Param5            int
	Param6            int
	Param7            int
	Param8            int
	ToHit             int
	LevToHit          int
	ResultFlags       int
	HitFlags          int
	HitClass          int
	HitShift          int
	SrcDam            int
	MinDam            int
	MinLevDam1        int
	MinLevDam2        int
	MinLevDam3        int
	MinLevDam4        int
	MinLevDam5        int
	MaxDam            int
	MaxLevDam1        int
	MaxLevDam2        int
	MaxLevDam3        int
	MaxLevDam4        int
	MaxLevDam5        int
	EMin              int
	EMinLev1          int
	EMinLev2          int
	EMinLev3          int
	EMinLev4          int
	EMinLev5          int
	EMax              int
	EMaxLev1          int
	EMaxLev2          int
	EMaxLev3          int
	EMaxLev4          int
	EMaxLev5          int
	ELen              int
	ELevLen1          int
	ELevLen2          int
	ELevLen3          int
	Aitype            int
	Aibonus           int
	CostMult          int
	CostAdd           int
	Prgstack          bool
	Decquant          bool
	Lob               bool
	Stsuccessonly     bool
	Stsounddelay      bool
	Weaponsnd         bool
	Warp              bool
	Immediate         bool
	Enhanceable       bool
	Noammo            bool
	Durability        bool
	UseAttackRate     bool
	TargetableOnly    bool
	SearchEnemyXY     bool
	SearchEnemyNear   bool
	SearchOpenXY      bool
	TargetCorpse      bool
	TargetPet         bool
	TargetAlly        bool
	TargetItem        bool
	AttackNoMana      bool
	TgtPlaceCheck     bool
	ItemCheckStart    bool
	ItemCltCheckStart bool
	Leftskill         bool
	Repeat            bool
	Nocostinstate     bool
	Usemanaondo       bool
	Interrupt         bool
	InTown            bool
	Aura              bool
	Periodic          bool
	Finishing         bool
	Passive           bool
	Progressive       bool
	General           bool
	Scroll            bool
	InGame            bool
	Kick              bool
}