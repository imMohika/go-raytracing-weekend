package main

type HittableList struct {
	objects []Hittable
}

func (l *HittableList) Add(object Hittable) {
	l.objects = append(l.objects, object)
}

func (l *HittableList) Hit(r Ray, t Interval) *HitRecord {
	var tempRec HitRecord
	closest := t.max
	for _, object := range l.objects {
		if rec := object.Hit(r, Interval{t.min, closest}); rec != nil {
			tempRec = *rec
			closest = rec.t
		}
	}

	return &tempRec
}
