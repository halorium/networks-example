# no shebang line as this is intended to be
# sourced:
#   source color/ansi.bash
#
# then used:
#   echo $"$Start$Yellow;$Underline$Stopbob$Reset"

Reset=0
Bold=1
Faint=2
Italic=3
Underline=4
BlinkSlow=5
BlinkRapid=6
Negative=7
Conceal=8
CrossedOut=9
PrimaryFont=10
FirstAlternateFont=11
SecondAlternateFont=12
ThirdAlternateFont=13
FourthAlternateFont=14
FifthAlternateFont=15
SixAlternateFont=16
SeventhAlternateFont=17
EighthAlternateFont=18
NinthAlternateFont=19
Fraktur=20
BoldOffDoubleUnderline=21
Normal=22
NotItalicNotFraktur=23
NoUnderline=24
BlinkOff=25

Positive=27
Reveal=28
NotCrossedOut=29

Black=30
Red=31
Green=32
Yellow=33
Blue=34
Purple=35
Cyan=36
White=37

BlackBackground=40
RedBackground=41
GreenBackground=42
YellowBackground=43
BlueBackground=44
PurpleBackground=45
CyanBackground=46
WhiteBackground=47

IntenseBlack=90
IntenseRed=91
IntenseGreen=92
IntenseYellow=93
IntenseBlue=94
IntensePurple=95
IntenseCyan=96
IntenseWhite=97

IntenseBlackBackground=100
IntenseRedBackground=101
IntenseGreenBackground=102
IntenseYellowBackground=103
IntenseBlueBackground=104
IntensePurpleBackground=105
IntenseCyanBackground=106
IntenseWhiteBackground=107

RGBForeground=38\;2\;
RGBBackground=48\;2\;

Start="\033"\[
Stop=m
Reset="\033"\[0m
