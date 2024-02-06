package stone1

type MetaTag uint16

const (
	// Name of the package
	Name MetaTag = iota + 1
	// Architecture is the architecture of the package.
	Architecture
	// Version is the version of the package.
	Version
	// Summary is the succint description of the package.
	Summary
	// Description is the description of the package.
	Description
	// Homepage is the homepage URL of the package.
	Homepage
	// SourceID is the ID of the source package used for grouping.
	SourceID
	// Depends is one dependency of the package.
	Depends
	// Provides is one capability, or the name, of the package.
	Provides
	// Conflicts is one capability, or name, conflicting with this package.
	Conflicts
	// Release is the release number of the package.
	Release
	// SPDX lists the SPDX license identifiers of the package.
	License
	// BuildRelease is the currently recorded build number of the package.
	BuildRelease
	// PackageURI is the URI of the package.
	PackageURI
	// PackageHash is the hash sum of the package.
	PackageHash
	// PackageSize is the size of the package.
	PackageSize
	// Depends is one build-time dependency of the package.
	BuildDepends
	// SourceURI is the URI of the source of the package.
	SourceURI
	// SourcePath is the relative path for the source within the upstream URI.
	SourcePath
	// SourceRef is the ref (or commit) of the upstream source.
	SourceRef
)

type MetaKind uint8

const (
	Int8 MetaKind = iota + 1
	Uint8
	Int16
	Uint16
	Int32
	Uint32
	Int64
	Uint64
	String
	Dependency
	Provider
)

type MetaValue struct {
	Kind  MetaKind
	Value any
}

func (mv MetaValue) Size() int {
	switch mv.Kind {
	case Int8, Uint8:
		return 1
	case Int16, Uint16:
		return 2
	case Int32, Uint32:
		return 4
	case Int64, Uint64:
		return 8
	case String:
		return len(mv.Value.(string))
	case Dependency, Provider:
		// TODO: should cast to a DependencyValue.
		return len(mv.Value.(string))
	default:
		panic("unknown MetaKind value")
	}
}

type DependencyKind uint8

const (
	PackageName DependencyKind = iota
	/// SharedLibary is a soname based dependency.
	SharedLibary
	/// PkgConfig is a pkgconfig `.pc` based dependency.
	PkgConfig
	/// Interpreter is a special interpreter (PT_INTERP/etc) to run the binaries.
	Interpreter
	/// CMake is a CMake module.
	CMake
	/// Python is Python module.
	Python
	/// Binary is a binary in /usr/bin.
	BinaryDep
	/// SystemBinary is a binary in /usr/sbin.
	SystemBinary
	/// PkgConfig32 is a emul32-compatible pkgconfig .pc dependency (contained in lib32/*.pc).
	PkgConfig32
)

func (d DependencyKind) String() string {
	switch d {
	case PackageName:
		return "name"
	case SharedLibary:
		return "soname"
	case PkgConfig:
		return "pkgconfig"
	case Interpreter:
		return "interpreter"
	case CMake:
		return "cmake"
	case Python:
		return "python"
	case BinaryDep:
		return "binary"
	case SystemBinary:
		return "sysbinary"
	case PkgConfig32:
		return "pkgconfig32"
	default:
		panic("unknown Dependency value")
	}
}

type MetaRecord struct {
	Tag   MetaTag
	Value MetaValue
}
