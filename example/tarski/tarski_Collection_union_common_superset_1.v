(* This file is generated by Why3's Coq driver *)
(* Beware! Only edit allowed sections below    *)
Require Import BuiltIn.
Require BuiltIn.
Require HighOrd.
Require set.Set.

Parameter comprehension: forall {a:Type} {a_WT:WhyType a}, (a -> bool) ->
  (set.Set.set a).

Axiom comprehension_def : forall {a:Type} {a_WT:WhyType a}, forall (p:(a ->
  bool)), forall (x:a), (set.Set.mem x (comprehension p)) <-> ((p x) = true).

Parameter fc: forall {a:Type} {a_WT:WhyType a}, (a -> bool) -> (set.Set.set
  a) -> (a -> bool).

Axiom fc_def : forall {a:Type} {a_WT:WhyType a}, forall (p:(a -> bool))
  (u:(set.Set.set a)) (x:a), (((fc p u) x) = true) <-> (((p x) = true) /\
  (set.Set.mem x u)).

(* Why3 assumption *)
Definition filter {a:Type} {a_WT:WhyType a} (p:(a -> bool)) (u:(set.Set.set
  a)): (set.Set.set a) := (comprehension (fc p u)).

Parameter fc1: forall {a:Type} {a_WT:WhyType a} {b:Type} {b_WT:WhyType b},
  (a -> b) -> (set.Set.set a) -> (b -> bool).

Axiom fc_def1 : forall {a:Type} {a_WT:WhyType a} {b:Type} {b_WT:WhyType b},
  forall (f:(a -> b)) (u:(set.Set.set a)) (y:b), (((fc1 f u) y) = true) <->
  exists x:a, (set.Set.mem x u) /\ (y = (f x)).

(* Why3 assumption *)
Definition map {a:Type} {a_WT:WhyType a} {b:Type} {b_WT:WhyType b} (f:(a ->
  b)) (u:(set.Set.set a)): (set.Set.set b) := (comprehension (fc1 f u)).

Axiom map_def : forall {a:Type} {a_WT:WhyType a} {b:Type} {b_WT:WhyType b},
  forall (f:(a -> b)) (u:(set.Set.set a)), forall (x:a), (set.Set.mem x u) ->
  (set.Set.mem (f x) (map f u)).

Parameter fc2: forall {a:Type} {a_WT:WhyType a}, (set.Set.set (set.Set.set
  a)) -> (a -> bool).

Axiom fc_def2 : forall {a:Type} {a_WT:WhyType a}, forall (fam:(set.Set.set
  (set.Set.set a))) (x:a), (((fc2 fam) x) = true) <-> forall (y:(set.Set.set
  a)), (set.Set.mem y fam) -> (set.Set.mem x y).

(* Why3 assumption *)
Definition intersect {a:Type} {a_WT:WhyType a} (fam:(set.Set.set (set.Set.set
  a))): (set.Set.set a) := (comprehension (fc2 fam)).

Axiom intersect_common_subset : forall {a:Type} {a_WT:WhyType a},
  forall (fam:(set.Set.set (set.Set.set a))), forall (x:(set.Set.set a)),
  (set.Set.mem x fam) -> (set.Set.subset (intersect fam) x).

Axiom intersect_greatest_common_subset : forall {a:Type} {a_WT:WhyType a},
  forall (fam:(set.Set.set (set.Set.set a))) (s:(set.Set.set a)),
  (forall (x:(set.Set.set a)), (set.Set.mem x fam) -> (set.Set.subset s
  x)) -> (set.Set.subset s (intersect fam)).

Parameter fc3: forall {a:Type} {a_WT:WhyType a}, (set.Set.set (set.Set.set
  a)) -> (a -> bool).

Axiom fc_def3 : forall {a:Type} {a_WT:WhyType a}, forall (fam:(set.Set.set
  (set.Set.set a))) (x:a), (((fc3 fam) x) = true) <-> exists y:(set.Set.set
  a), (set.Set.mem y fam) /\ (set.Set.mem x y).

(* Why3 assumption *)
Definition union {a:Type} {a_WT:WhyType a} (fam:(set.Set.set (set.Set.set
  a))): (set.Set.set a) := (comprehension (fc3 fam)).

Parameter fc4: forall {a:Type} {a_WT:WhyType a}, (set.Set.set (set.Set.set
  a)) -> (a -> bool).

Axiom fc_def4 : forall {a:Type} {a_WT:WhyType a}, forall (fam:(set.Set.set
  (set.Set.set a))) (x:a), (((fc4 fam) x) = true) <-> exists y:(set.Set.set
  a), (set.Set.mem y fam) /\ (set.Set.mem x y).

(* Why3 goal *)
Theorem union_common_superset : forall {a:Type} {a_WT:WhyType a},
  forall (fam:(set.Set.set (set.Set.set a))), forall (x:(set.Set.set a)),
  (set.Set.mem x fam) -> forall (x1:a), (set.Set.mem x1 x) -> (set.Set.mem x1
  (comprehension (fc4 fam))).
intros a a_WT fam x h1 x1 h2.
rewrite comprehension_def.
rewrite fc_def4.
exists x.
split.
exact h1.
exact h2.
Qed.
