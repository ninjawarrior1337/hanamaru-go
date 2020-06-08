package voice

type Queue struct {
	audioFiles []Playable
}

func (q *Queue) Push(p Playable) {
	q.audioFiles = append(q.audioFiles, p)
}

func (q *Queue) Length() int {
	return len(q.audioFiles)
}

func (q *Queue) Pop() Playable {
	element := q.audioFiles[0]
	q.audioFiles = q.audioFiles[1:]
	return element
}
