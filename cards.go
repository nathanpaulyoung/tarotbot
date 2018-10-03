package main

import (
	"fmt"
	"strings"
)

type card struct {
	name     string
	abbrev   string
	upright  string
	inverted string
	link     string
}

func getCardByAbbreviation(abbrev string) (card, bool, error) {
	var inverted bool

	if strings.HasSuffix(abbrev, "i") {
		inverted = true
		abbrev = strings.TrimSuffix(abbrev, "i")
	}

	if strings.HasPrefix(abbrev, "i") {
		inverted = true
		abbrev = strings.TrimPrefix(abbrev, "i")
	}

	if strings.HasSuffix(abbrev, "r") {
		inverted = true
		abbrev = strings.TrimSuffix(abbrev, "r")
	}

	if strings.HasPrefix(abbrev, "r") {
		inverted = true
		abbrev = strings.TrimPrefix(abbrev, "r")
	}

	switch abbrev {
	case aceOfPentacles.abbrev:
		return aceOfPentacles, inverted, nil
	case twoOfPentacles.abbrev:
		return twoOfPentacles, inverted, nil
	case threeOfPentacles.abbrev:
		return threeOfPentacles, inverted, nil
	case fourOfPentacles.abbrev:
		return fourOfPentacles, inverted, nil
	case fiveOfPentacles.abbrev:
		return fiveOfPentacles, inverted, nil
	case sixOfPentacles.abbrev:
		return sixOfPentacles, inverted, nil
	case sevenOfPentacles.abbrev:
		return sevenOfPentacles, inverted, nil
	case eightOfPentacles.abbrev:
		return eightOfPentacles, inverted, nil
	case nineOfPentacles.abbrev:
		return nineOfPentacles, inverted, nil
	case tenOfPentacles.abbrev:
		return tenOfPentacles, inverted, nil
	case pageOfPentacles.abbrev:
		return pageOfPentacles, inverted, nil
	case knightOfPentacles.abbrev:
		return knightOfPentacles, inverted, nil
	case queenOfPentacles.abbrev:
		return queenOfPentacles, inverted, nil
	case kingOfPentacles.abbrev:
		return kingOfPentacles, inverted, nil
	case aceOfSwords.abbrev:
		return aceOfSwords, inverted, nil
	case twoOfSwords.abbrev:
		return twoOfSwords, inverted, nil
	case threeOfSwords.abbrev:
		return threeOfSwords, inverted, nil
	case fourOfSwords.abbrev:
		return fourOfSwords, inverted, nil
	case fiveOfSwords.abbrev:
		return fiveOfSwords, inverted, nil
	case sixOfSwords.abbrev:
		return sixOfSwords, inverted, nil
	case sevenOfSwords.abbrev:
		return sevenOfSwords, inverted, nil
	case eightOfSwords.abbrev:
		return eightOfSwords, inverted, nil
	case nineOfSwords.abbrev:
		return nineOfSwords, inverted, nil
	case tenOfSwords.abbrev:
		return tenOfSwords, inverted, nil
	case pageOfSwords.abbrev:
		return pageOfSwords, inverted, nil
	case knightOfSwords.abbrev:
		return knightOfSwords, inverted, nil
	case queenOfSwords.abbrev:
		return queenOfSwords, inverted, nil
	case kingOfSwords.abbrev:
		return kingOfSwords, inverted, nil
	case aceOfCups.abbrev:
		return aceOfCups, inverted, nil
	case twoOfCups.abbrev:
		return twoOfCups, inverted, nil
	case threeOfCups.abbrev:
		return threeOfCups, inverted, nil
	case fourOfCups.abbrev:
		return fourOfCups, inverted, nil
	case fiveOfCups.abbrev:
		return fiveOfCups, inverted, nil
	case sixOfCups.abbrev:
		return sixOfCups, inverted, nil
	case sevenOfCups.abbrev:
		return sevenOfCups, inverted, nil
	case eightOfCups.abbrev:
		return eightOfCups, inverted, nil
	case nineOfCups.abbrev:
		return nineOfCups, inverted, nil
	case tenOfCups.abbrev:
		return tenOfCups, inverted, nil
	case pageOfCups.abbrev:
		return pageOfCups, inverted, nil
	case knightOfCups.abbrev:
		return knightOfCups, inverted, nil
	case queenOfCups.abbrev:
		return queenOfCups, inverted, nil
	case kingOfCups.abbrev:
		return kingOfCups, inverted, nil
	case aceOfWands.abbrev:
		return aceOfWands, inverted, nil
	case twoOfWands.abbrev:
		return twoOfWands, inverted, nil
	case threeOfWands.abbrev:
		return threeOfWands, inverted, nil
	case fourOfWands.abbrev:
		return fourOfWands, inverted, nil
	case fiveOfWands.abbrev:
		return fiveOfWands, inverted, nil
	case sixOfWands.abbrev:
		return sixOfWands, inverted, nil
	case sevenOfWands.abbrev:
		return sevenOfWands, inverted, nil
	case eightOfWands.abbrev:
		return eightOfWands, inverted, nil
	case nineOfWands.abbrev:
		return nineOfWands, inverted, nil
	case tenOfWands.abbrev:
		return tenOfWands, inverted, nil
	case pageOfWands.abbrev:
		return pageOfWands, inverted, nil
	case knightOfWands.abbrev:
		return knightOfWands, inverted, nil
	case queenOfWands.abbrev:
		return queenOfWands, inverted, nil
	case kingOfWands.abbrev:
		return kingOfWands, inverted, nil
	case theFool.abbrev:
		return theFool, inverted, nil
	case theMagician.abbrev:
		return theMagician, inverted, nil
	case theHighPriestess.abbrev:
		return theHighPriestess, inverted, nil
	case theEmpress.abbrev:
		return theEmpress, inverted, nil
	case theEmperor.abbrev:
		return theEmperor, inverted, nil
	case theHierophant.abbrev:
		return theHierophant, inverted, nil
	case theLovers.abbrev:
		return theLovers, inverted, nil
	case theChariot.abbrev:
		return theChariot, inverted, nil
	case strength.abbrev:
		return strength, inverted, nil
	case theHermit.abbrev:
		return theHermit, inverted, nil
	case wheelOfFortune.abbrev:
		return wheelOfFortune, inverted, nil
	case justice.abbrev:
		return justice, inverted, nil
	case theHangedMan.abbrev:
		return theHangedMan, inverted, nil
	case death.abbrev:
		return death, inverted, nil
	case temperance.abbrev:
		return temperance, inverted, nil
	case theDevil.abbrev:
		return theDevil, inverted, nil
	case theTower.abbrev:
		return theTower, inverted, nil
	case theStars.abbrev:
		return theStars, inverted, nil
	case theMoon.abbrev:
		return theMoon, inverted, nil
	case theSun.abbrev:
		return theSun, inverted, nil
	case judgement.abbrev:
		return judgement, inverted, nil
	case theWorld.abbrev:
		return theWorld, inverted, nil
	default:
		return card{"", "", "", "", ""}, inverted, fmt.Errorf("%v is an invalid card abbreviation", abbrev)
	}
}

var aceOfPentacles = card{
	name:     "Ace of Pentacles",
	abbrev:   "AP",
	upright:  "manifestation, new financial opportunity, prosperity",
	inverted: "lost opportunity, lack of planning and foresight",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/ace-of-pentacles/",
}

var twoOfPentacles = card{
	name:     "Two of Pentacles",
	abbrev:   "2P",
	upright:  "balance, adaptability, time management, prioritization",
	inverted: "disorganization, financial disarray",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/two-of-pentacles/",
}

var threeOfPentacles = card{
	name:     "Three of Pentacles",
	abbrev:   "3P",
	upright:  "teamwork, initial fulfilment, collaboration, learning",
	inverted: "lack of teamwork, disregard for skills",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/three-of-pentacles/",
}

var fourOfPentacles = card{
	name:     "Four of Pentacles",
	abbrev:   "4P",
	upright:  "control, stability, security, possession, conservatism",
	inverted: "Greed, materialism, self-protection",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/four-of-pentacles/",
}

var fiveOfPentacles = card{
	name:     "Five of Pentacles",
	abbrev:   "5P",
	upright:  "isolation, insecurity, worry, financial loss, poverty",
	inverted: "recovery from financial loss, spiritual poverty",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/five-of-pentacles/",
}

var sixOfPentacles = card{
	name:     "Six of Pentacles",
	abbrev:   "6P",
	upright:  "generosity, charity, giving, prosperity, sharing wealth",
	inverted: "debt, selfishness, one-sided charity",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/six-of-pentacles/",
}

var sevenOfPentacles = card{
	name:     "Seven of Pentacles",
	abbrev:   "7P",
	upright:  "vision, perseverance, profit, reward, investment",
	inverted: "lack of long-term vision, limited success or reward",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/seven-of-pentacles/",
}

var eightOfPentacles = card{
	name:     "Eight of Pentacles",
	abbrev:   "8P",
	upright:  "apprenticeship, education, quality, engagement",
	inverted: "perfectionism, lacking ambition or focus",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/eight-of-pentacles/",
}

var nineOfPentacles = card{
	name:     "Nine of Pentacles",
	abbrev:   "9P",
	upright:  "gratitude, luxury, self-sufficiency, culmination",
	inverted: "over-investment in work, financial setbacks",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/nine-of-pentacles/",
}

var tenOfPentacles = card{
	name:     "Ten of Pentacles",
	abbrev:   "10P",
	upright:  "wealth, inheritance, family, establishment, retirement",
	inverted: "financial failure, loneliness, loss",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/ten-of-pentacles/",
}

var pageOfPentacles = card{
	name:     "Page of Pentacles",
	abbrev:   "PP",
	upright:  "manifestation, financial opportunity, new job",
	inverted: "lack of progress and planning, short-term focus",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/page-of-pentacles/",
}

var knightOfPentacles = card{
	name:     "Knight of Pentacles",
	abbrev:   "NP",
	upright:  "efficiency, routine, conservatism, methodical",
	inverted: "Laziness, boredom, feeling stuck",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/knight-of-pentacles/",
}

var queenOfPentacles = card{
	name:     "Queen of Pentacles",
	abbrev:   "QP",
	upright:  "practical, homely, motherly, down-to-earth, security",
	inverted: "imbalance in work or with family commitments",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/queen-of-pentacles/",
}

var kingOfPentacles = card{
	name:     "King of Pentacles",
	abbrev:   "KP",
	upright:  "security, control, power, discipline, abundance",
	inverted: "authoritative, domineering, controlling",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-pentacles/king-of-pentacles/",
}

var aceOfSwords = card{
	name:     "Ace of Swords",
	abbrev:   "AS",
	upright:  "raw power, victory, break-throughs, mental clarity",
	inverted: "confusion, chaos, lack of clarity",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/ace-of-swords/",
}

var twoOfSwords = card{
	name:     "Two of Swords",
	abbrev:   "2S",
	upright:  "indecision, choices, truce, stalemate, blocked emotions",
	inverted: "indecision, confusion, information overload",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/two-of-swords/",
}

var threeOfSwords = card{
	name:     "Three of Swords",
	abbrev:   "3S",
	upright:  "painful separation, sorrow, heartbreak, grief, rejection",
	inverted: "releasing pain, optimism, forgiveness",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/three-of-swords/",
}

var fourOfSwords = card{
	name:     "Four of Swords",
	abbrev:   "4S",
	upright:  "contemplation, recuperation, passivity, relaxation, rest",
	inverted: "restlessness, burn-out, lack of progress",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/four-of-swords/",
}

var fiveOfSwords = card{
	name:     "Five of Swords",
	abbrev:   "5S",
	upright:  "conflict, tension, loss, defeat, win at all costs, betrayal",
	inverted: "open to change, past resentment",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/five-of-swords/",
}

var sixOfSwords = card{
	name:     "Six of Swords",
	abbrev:   "6S",
	upright:  "regretful but necessary transition, rite of passage",
	inverted: "cannot move on, carrying baggage",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/six-of-swords/",
}

var sevenOfSwords = card{
	name:     "Seven of Swords",
	abbrev:   "7S",
	upright:  "betrayal, deception, getting away with something, stealth",
	inverted: "mental challenges, breaking free",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/seven-of-swords/",
}

var eightOfSwords = card{
	name:     "Eight of Swords",
	abbrev:   "8S",
	upright:  "isolation, self-imposed restriction, imprisonment",
	inverted: "open to new perspectives, release",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/eight-of-swords/",
}

var nineOfSwords = card{
	name:     "Nine of Swords",
	abbrev:   "9S",
	upright:  "depression, nightmares, intense anxiety, despair",
	inverted: "hopelessness, severe depression, torment",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/nine-of-swords/",
}

var tenOfSwords = card{
	name:     "Ten of Swords",
	abbrev:   "10S",
	upright:  "back-stabbed, defeat, crisis, betrayal, endings, loss",
	inverted: "recovery, regeneration, fear of ruin, inevitable end",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/ten-of-swords/",
}

var pageOfSwords = card{
	name:     "Page of Swords",
	abbrev:   "PS",
	upright:  "talkative, curious, mentally restless, energetic",
	inverted: "all talk and no action, haste, undelivered promises",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/page-of-swords/",
}

var knightOfSwords = card{
	name:     "Knight of Swords",
	abbrev:   "NS",
	upright:  "opinionated, hasty, action-oriented, communicative",
	inverted: "scattered thought, disregard for consequences",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/knight-of-swords/",
}

var queenOfSwords = card{
	name:     "Queen of Swords",
	abbrev:   "QS",
	upright:  "quick thinker, organized, perceptive, independent",
	inverted: "overly-emotional, bitchy, cold-hearted",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/queen-of-swords/",
}

var kingOfSwords = card{
	name:     "King of Swords",
	abbrev:   "KS",
	upright:  "clear thinking, intellectual power, authority, truth",
	inverted: "manipulative, tyrannical, abusive",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-swords/king-of-swords/",
}

var aceOfCups = card{
	name:     "Ace of Cups",
	abbrev:   "AC",
	upright:  "love, new relationships, compassion, creativity",
	inverted: "self-love, intuition, repressed emotions",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/ace-of-cups/",
}

var twoOfCups = card{
	name:     "Two of Cups",
	abbrev:   "2C",
	upright:  "unified love, partnership, mutual attraction",
	inverted: "self-love, break-ups, disharmony, distrust",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/two-of-cups/",
}

var threeOfCups = card{
	name:     "Three of Cups",
	abbrev:   "3C",
	upright:  "celebration, friendship, creativity, collaborations",
	inverted: "independence, alone time, hardcore partying, three's a crowd",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/three-of-cups/",
}

var fourOfCups = card{
	name:     "Four of Cups",
	abbrev:   "4C",
	upright:  "meditation, contemplation, apathy, re-evaluation",
	inverted: "boredom, missed opportunity, being aloof",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/four-of-cups/",
}

var fiveOfCups = card{
	name:     "Five of Cups",
	abbrev:   "5C",
	upright:  "loss, regret, disappointment, despair, bereavement",
	inverted: "moving on, acceptance, forgiveness",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/five-of-cups/",
}

var sixOfCups = card{
	name:     "Six of Cups",
	abbrev:   "6C",
	upright:  "reunion, nostalgia, childhood memories, innocence",
	inverted: "stuck in the past, naivety, unrealistic",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/six-of-cups/",
}

var sevenOfCups = card{
	name:     "Seven of Cups",
	abbrev:   "7C",
	upright:  "fantasy, illusion, wishful thinking, choices, imagination",
	inverted: "temptation, illusion, diversionary tactics",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/seven-of-cups/",
}

var eightOfCups = card{
	name:     "Eight of Cups",
	abbrev:   "8C",
	upright:  "escapism, disappointment, abandonment, withdrawal",
	inverted: "hopelessness, aimless drifting, walking away",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/eight-of-cups/",
}

var nineOfCups = card{
	name:     "Nine of Cups",
	abbrev:   "9C",
	upright:  "wishes fulfilled, comfort, happiness, satisfaction",
	inverted: "greed, dissatisfaction, materialism",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/nine-of-cups/",
}

var tenOfCups = card{
	name:     "Ten of Cups",
	abbrev:   "10C",
	upright:  "harmony, marriage, happiness, alignment",
	inverted: "misalignment of values, broken home or marriage",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/ten-of-cups/",
}

var pageOfCups = card{
	name:     "Page of Cups",
	abbrev:   "PC",
	upright:  "a messenger, creative beginnings, synchronicity",
	inverted: "emotional immaturity, creative block",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/page-of-cups/",
}

var knightOfCups = card{
	name:     "Knight of Cups",
	abbrev:   "NC",
	upright:  "romance, charm, knight in shining armor, imagination",
	inverted: "unrealistic, jealousy, moodiness",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/knight-of-cups/",
}

var queenOfCups = card{
	name:     "Queen of Cups",
	abbrev:   "QC",
	upright:  "emotional security, calm, intuitive, compassionate",
	inverted: "emotional insecurity, co-dependency",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/queen-of-cups/",
}

var kingOfCups = card{
	name:     "King of Cups",
	abbrev:   "KC",
	upright:  "emotional balance and control, generosity",
	inverted: "emotional manipulation, moodiness, volatility",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-cups/king-of-cups/",
}

var aceOfWands = card{
	name:     "Ace of Wands",
	abbrev:   "AW",
	upright:  "inspiration, power, creation, beginnings, potential",
	inverted: "delays, lack of motivation, weighed down",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/ace-of-wands/",
}

var twoOfWands = card{
	name:     "Two of Wands",
	abbrev:   "2W",
	upright:  "future planning, progress, decisions, discovery",
	inverted: "fear of unknown, lack of planning",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/two-of-wands/",
}

var threeOfWands = card{
	name:     "Three of Wands",
	abbrev:   "3W",
	upright:  "preparation, foresight, enterprise, expansion",
	inverted: "lack of foresight, delays, obstacles to long-term goals",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/three-of-wands/",
}

var fourOfWands = card{
	name:     "Four of Wands",
	abbrev:   "4W",
	upright:  "celebration, harmony, marriage, home, community",
	inverted: "breakdown in communication, transition",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/four-of-wands/",
}

var fiveOfWands = card{
	name:     "Five of Wands",
	abbrev:   "5W",
	upright:  "disagreement, competition, strife, tension, conflict",
	inverted: "conflict avoidance, increased focus on goals",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/five-of-wands/",
}

var sixOfWands = card{
	name:     "Six of Wands",
	abbrev:   "6W",
	upright:  "public recognition, victory, progress, self-confidence",
	inverted: "egotism, disrepute, lack of confidence, fall from grace",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/six-of-wands/",
}

var sevenOfWands = card{
	name:     "Seven of Wands",
	abbrev:   "7W",
	upright:  "challenge, competition, perseverance",
	inverted: "giving up, overwhelmed, overly protective",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/seven-of-wands/",
}

var eightOfWands = card{
	name:     "Eight of Wands",
	abbrev:   "8W",
	upright:  "speed, action, air travel, movement, swift change",
	inverted: "delays, frustration, holding off",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/eight-of-wands/",
}

var nineOfWands = card{
	name:     "Nine of Wands",
	abbrev:   "9W",
	upright:  "courage, persistence, test of faith, resilience",
	inverted: "on edge, defensive, hesitant, paranoia",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/nine-of-wands/",
}

var tenOfWands = card{
	name:     "Ten of Wands",
	abbrev:   "10W",
	upright:  "burden, responsibility, hard work, stress, achievement",
	inverted: "taking on too much, avoiding responsibility",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/ten-of-wands/",
}

var pageOfWands = card{
	name:     "Page of Wands",
	abbrev:   "PW",
	upright:  "enthusiasm, exploration, discovery, free spirit",
	inverted: "Setbacks to new ideas, pessimism, lack of direction",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/page-of-wands/",
}

var knightOfWands = card{
	name:     "Knight of Wands",
	abbrev:   "NW",
	upright:  "energy, passion, lust, action, adventure, impulsiveness",
	inverted: "haste, scattered energy, delays, frustration",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/knight-of-wands/",
}

var queenOfWands = card{
	name:     "Queen of Wands",
	abbrev:   "QW",
	upright:  "exuberance, warmth, vibrancy, determination",
	inverted: "shrinking violet, aggressive, demanding",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/queen-of-wands/",
}

var kingOfWands = card{
	name:     "King of Wands",
	abbrev:   "KW",
	upright:  "natural-born leader, vision, entrepreneur, honour",
	inverted: "impulsiveness, haste, ruthless, high expectations",
	link:     "https://www.biddytarot.com/tarot-card-meanings/minor-arcana/suit-of-wands/king-of-wands/",
}

var theFool = card{
	name:     "The Fool",
	abbrev:   "0",
	upright:  "beginnings, innocence, spontaneity, a free spirit",
	inverted: "naivety, foolishness, recklessness, risk-taking",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/fool/",
}

var theMagician = card{
	name:     "The Magician",
	abbrev:   "I",
	upright:  "manifestation, resourcefulness, power, inspired action",
	inverted: "manipulation, poor planning, untapped talents",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/magician/",
}

var theHighPriestess = card{
	name:     "The High Priestess",
	abbrev:   "II",
	upright:  "intuition, sacred knowledge, divine feminine, the subconscious mind.",
	inverted: "secrets, disconnected from intuition, withdrawal and silence",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/high-priestess/",
}

var theEmpress = card{
	name:     "The Empress",
	abbrev:   "III",
	upright:  "femininity, beauty, nature, nurturing, abundance",
	inverted: "creative block, dependence on others",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/empress/",
}

var theEmperor = card{
	name:     "The Emperor",
	abbrev:   "IV",
	upright:  "authority, establishment, structure, a father figure",
	inverted: "domination, excessive control, inflexibility, self-discipline",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/emperor/",
}

var theHierophant = card{
	name:     "The Hierophant",
	abbrev:   "V",
	upright:  "spiritual wisdom, religious beliefs, conformity, tradition, institutions",
	inverted: "personal beliefs, freedom, challenging the status quo",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/hierophant/",
}

var theLovers = card{
	name:     "The Lovers",
	abbrev:   "VI",
	upright:  "love, harmony, relationships, values alignment, choices",
	inverted: "self-love, disharmony, imbalance, misalignment of values",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/lovers/",
}

var theChariot = card{
	name:     "The Chariot",
	abbrev:   "VII",
	upright:  "control, willpower, success, action, determination",
	inverted: "self-discipline, opposition, lack of direction",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/chariot/",
}

var strength = card{
	name:     "Strength",
	abbrev:   "VIII",
	upright:  "strength, courage, persuasion, influence, compassion",
	inverted: "inner strength, self-doubt, low energy, raw emotion",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/strength/",
}

var theHermit = card{
	name:     "The Hermit",
	abbrev:   "IX",
	upright:  "soul-searching, introspection, being alone, inner guidance",
	inverted: "isolation, loneliness, withdrawal",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/hermit/",
}

var wheelOfFortune = card{
	name:     "Wheel of Fortune",
	abbrev:   "X",
	upright:  "good luck, karma, life cycles, destiny, a turning point",
	inverted: "bad luck, resistance to change, breaking cycles",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/wheel-of-fortune/",
}

var justice = card{
	name:     "Justice",
	abbrev:   "XI",
	upright:  "justice, fairness, truth, cause and effect, law",
	inverted: "unfairness, lack of accountability, dishonesty",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/justice/",
}

var theHangedMan = card{
	name:     "The Hanged Man",
	abbrev:   "XII",
	upright:  "pause, surrender, letting go, new perspectives",
	inverted: "delays, resistance, stalling, indecision",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/hanged-man/",
}

var death = card{
	name:     "Death",
	abbrev:   "XIII",
	upright:  "endings, change, transformation, transition",
	inverted: "Resistance to change, personal transformation, inner purging",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/death/",
}

var temperance = card{
	name:     "Temperance",
	abbrev:   "XIV",
	upright:  "balance, moderation, patience, purpose",
	inverted: "imbalance, excess, self-healing, re-alignment",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/temperance/",
}

var theDevil = card{
	name:     "The Devil",
	abbrev:   "XV",
	upright:  "shadow self, attachment, addiction, restriction, sexuality",
	inverted: "releasing limiting beliefs, exploring dark thoughts, detachment",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/devil/",
}

var theTower = card{
	name:     "The Tower",
	abbrev:   "XVI",
	upright:  "sudden change, upheaval, chaos, revelation, awakening",
	inverted: "personal transformation, fear of change, averting disaster",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/tower/",
}

var theStars = card{
	name:     "The Stars",
	abbrev:   "XVII",
	upright:  "hope, faith, purpose, renewal, spirituality",
	inverted: "lack of faith, despair, self-trust, disconnection",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/star/",
}

var theMoon = card{
	name:     "The Moon",
	abbrev:   "XVIII",
	upright:  "illusion, fear, anxiety, subconscious, intuition",
	inverted: "release of fear, repressed emotion, inner confusion",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/moon/",
}

var theSun = card{
	name:     "The Sun",
	abbrev:   "XIX",
	upright:  "positivity, fun, warmth, success, vitality",
	inverted: "inner child, feeling down, overly optimistic",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/sun/",
}

var judgement = card{
	name:     "Judgement",
	abbrev:   "XX",
	upright:  "judgement, rebirth, inner calling, absolution",
	inverted: "self-doubt, inner critic, ignoring the call",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/judgement/",
}

var theWorld = card{
	name:     "The World",
	abbrev:   "XXI",
	upright:  "completion, integration, accomplishment, travel",
	inverted: "seeking personal closure, short-cuts, delays",
	link:     "https://www.biddytarot.com/tarot-card-meanings/major-arcana/world/",
}
