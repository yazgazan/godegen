package utils

import (
	"iter"
	"strings"
	"unicode"
)

type Strings struct{}

func (Strings) Clone(s string) string {
	return strings.Clone(s)
}

func (Strings) Compare(a, b string) int {
	return strings.Compare(a, b)
}

func (Strings) Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func (Strings) ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

func (Strings) ContainsFunc(s string, f func(rune) bool) bool {
	return strings.ContainsFunc(s, f)
}

func (Strings) ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

func (Strings) Count(s, substr string) int {
	return strings.Count(s, substr)
}

func (Strings) Cut(s, sep string) (before, after string, found bool) {
	return strings.Cut(s, sep)
}

func (Strings) CutPrefix(s, prefix string) (after string, found bool) {
	return strings.CutPrefix(s, prefix)
}

func (Strings) CutSuffix(s, suffix string) (before string, found bool) {
	return strings.CutSuffix(s, suffix)
}

func (Strings) EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

func (Strings) Fields(s string) []string {
	return strings.Fields(s)
}

func (Strings) FieldsFunc(s string, f func(rune) bool) []string {
	return strings.FieldsFunc(s, f)
}

func (Strings) FieldsFuncSeq(s string, f func(rune) bool) iter.Seq[string] {
	return strings.FieldsFuncSeq(s, f)
}

func (Strings) FieldsSeq(s string) iter.Seq[string] {
	return strings.FieldsSeq(s)
}

func (Strings) HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func (Strings) HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func (Strings) Index(s, substr string) int {
	return strings.Index(s, substr)
}

func (Strings) IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}

func (Strings) IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}

func (Strings) IndexFunc(s string, f func(rune) bool) int {
	return strings.IndexFunc(s, f)
}

func (Strings) IndexRune(s string, r rune) int {
	return strings.IndexRune(s, r)
}

func (Strings) Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

func (Strings) LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

func (Strings) LastIndexAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}

func (Strings) LastIndexByte(s string, c byte) int {
	return strings.LastIndexByte(s, c)
}

func (Strings) LastIndexFunc(s string, f func(rune) bool) int {
	return strings.LastIndexFunc(s, f)
}

func (Strings) Lines(s string) iter.Seq[string] {
	return strings.Lines(s)
}

func (Strings) Map(mapping func(rune) rune, s string) string {
	return strings.Map(mapping, s)
}

func (Strings) Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

func (Strings) Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func (Strings) ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func (Strings) Split(s, sep string) []string {
	return strings.Split(s, sep)
}

func (Strings) SplitAfter(s, sep string) []string {
	return strings.SplitAfter(s, sep)
}

func (Strings) SplitAfterN(s, sep string, n int) []string {
	return strings.SplitAfterN(s, sep, n)
}

func (Strings) SplitAfterSeq(s, sep string) iter.Seq[string] {
	return strings.SplitAfterSeq(s, sep)
}

func (Strings) SplitN(s, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}

func (Strings) SplitSeq(s, sep string) iter.Seq[string] {
	return strings.SplitSeq(s, sep)
}

func (Strings) Title(s string) string {
	return strings.Title(s)
}

func (Strings) ToLower(s string) string {
	return strings.ToLower(s)
}

func (Strings) ToLowerSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToLowerSpecial(c, s)
}

func (Strings) ToTitle(s string) string {
	return strings.ToTitle(s)
}

func (Strings) ToTitleSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToTitleSpecial(c, s)
}

func (Strings) ToUpper(s string) string {
	return strings.ToUpper(s)
}

func (Strings) ToUpperSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToUpperSpecial(c, s)
}

func (Strings) ToValidUTF8(s, replacement string) string {
	return strings.ToValidUTF8(s, replacement)
}

func (Strings) Trim(s, cutset string) string {
	return strings.Trim(s, cutset)
}

func (Strings) TrimFunc(s string, f func(rune) bool) string {
	return strings.TrimFunc(s, f)
}

func (Strings) TrimLeft(s, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

func (Strings) TrimLeftFunc(s string, f func(rune) bool) string {
	return strings.TrimLeftFunc(s, f)
}

func (Strings) TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func (Strings) TrimRight(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}

func (Strings) TrimRightFunc(s string, f func(rune) bool) string {
	return strings.TrimRightFunc(s, f)
}

func (Strings) TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func (Strings) TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}
