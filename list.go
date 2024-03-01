package main

type HittableList struct {
	objects []Hittable
}

func (l *HittableList) Add(object Hittable) {
	l.objects = append(l.objects, object)
}

func (l *HittableList) Hit(r Ray, t Interval) (bool, HitRecord) {
	var tempRec HitRecord
	closest := t.max
	hit := false
	for _, object := range l.objects {
		if rec := object.Hit(r, Interval{t.min, closest}); rec != nil {
			tempRec = *rec
			closest = rec.t
			hit = true
		}
	}

	return hit, tempRec
}
