package ecs

// A System implements logic for processing entities possessing components of
// the same aspects as the system. A System should iterate over its Entities on
// `Update`, in any way suitable for the current implementation.
//
// By convention, systems provide an Add method for adding entities and their
// associated components to the system; e.g.
//
//	Add(basic *ecs.BasicEntity, collision *CollisionComponent, space *SpaceComponent)
type System interface {
	// Update updates the system. It is invoked by the engine once every frame,
	// with dt being the duration since the previous update.
	Update(dt float64)

	// Remove removes the given entity from the system.
	Remove(e BasicEntity)
}

// SystemAddByInterface is a system that also implements the AddByInterface method
type SystemAddByInterface interface {
	System

	// AddByInterface allows you to automatically add entities based on the
	// interfaces that the entity implements. It should add the entity passed
	// as o to the system after casting it to the correct interface.
	AddByInterface(o Identifier)
}

// Prioritize specifies the priority of systems.
type Prioritize interface {
	// Priority indicates the order in which Systems should be executed per
	// iteration, higher meaning sooner. The default priority is 0.
	Priority() int
}

// Initializer provides initialization of systems.
type Initializer interface {
	// New initializes the given System, and may be used to initialize some
	// values beforehand, like storing a reference to the World.
	New(*World)
}

// systems implements a sortable list of `System`. It is indexed on
// `System.Priority()`.
type systems []System

func (s systems) Len() int {
	return len(s)
}

func (s systems) Less(i, j int) bool {
	var _prior1, _prior2 int

	if prior1, ok := s[i].(Prioritize); ok {
		_prior1 = prior1.Priority()
	}
	if prior2, ok := s[j].(Prioritize); ok {
		_prior2 = prior2.Priority()
	}

	return _prior1 > _prior2
}

func (s systems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
