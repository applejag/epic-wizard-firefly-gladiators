(import "net" "get_me"
 (func $firefly/firefly/internal/ffi.get_me (result i32)))
(data  (memory $moonbit.memory) (offset (i32.const 10000)) "\FF\FF\FF\FF\00\00\00@\00\00\00\00\00\00\00\00\FF\FF\FF\FF\00\00\00`\00\00\00\00\00\00\00\00")
(memory $moonbit.memory 1)
(export "memory" (memory $moonbit.memory))
(global $tlsf/ROOT
 (mut i32)
 (i32.const 0)
)
(func $tlsf/searchBlock (param $0 i32) (param $1 i32) (result i32)
 (local $2 i32)
 (if (result i32)
  (local.tee $1
   (i32.and
    (i32.load offset=4
     (i32.add
      (local.get $0)
      (i32.shl
       (local.tee $2
        (if (result i32)
         (i32.lt_u
          (local.get $1)
          (i32.const 256))
         (then
          (local.set $1
           (i32.shr_u
            (local.get $1)
            (i32.const 4)))
          (i32.const 0))
         (else
          (if
           (i32.lt_u
            (local.get $1)
            (i32.const 536870910))
           (then
            (local.set $1
             (i32.sub
              (i32.add
               (local.get $1)
               (i32.shl
                (i32.const 1)
                (i32.sub
                 (i32.const 27)
                 (i32.clz
                  (local.get $1)))))
              (i32.const 1))))
           (else))
          (local.set $2
           (i32.sub
            (i32.const 31)
            (i32.clz
             (local.get $1))))
          (local.set $1
           (i32.xor
            (i32.shr_u
             (local.get $1)
             (i32.sub
              (local.get $2)
              (i32.const 4)))
            (i32.const 16)))
          (i32.sub
           (local.get $2)
           (i32.const 7)))))
       (i32.const 2))))
    (i32.shl
     (i32.const -1)
     (local.get $1))))
  (then
   (i32.load offset=96
    (i32.add
     (local.get $0)
     (i32.shl
      (i32.add
       (i32.ctz
        (local.get $1))
       (i32.shl
        (local.get $2)
        (i32.const 4)))
      (i32.const 2)))))
  (else
   (if (result i32)
    (local.tee $1
     (i32.and
      (i32.load
       (local.get $0))
      (i32.shl
       (i32.const -1)
       (i32.add
        (local.get $2)
        (i32.const 1)))))
    (then
     (i32.load offset=96
      (i32.add
       (local.get $0)
       (i32.shl
        (i32.add
         (i32.ctz
          (i32.load offset=4
           (i32.add
            (local.get $0)
            (i32.shl
             (local.tee $1
              (i32.ctz
               (local.get $1)))
             (i32.const 2)))))
         (i32.shl
          (local.get $1)
          (i32.const 4)))
        (i32.const 2)))))
    (else
     (i32.const 0))))))
(func $tlsf/removeBlock (param $0 i32) (param $1 i32)
 (local $2 i32)
 (local $3 i32)
 (local $4 i32)
 (local $5 i32)
 (local.set $5
  (if (result i32)
   (i32.lt_u
    (local.tee $2
     (i32.and
      (i32.load
       (local.get $1))
      (i32.const -4)))
    (i32.const 256))
   (then
    (local.set $3
     (i32.shr_u
      (local.get $2)
      (i32.const 4)))
    (i32.const 0))
   (else
    (local.set $2
     (i32.sub
      (i32.const 31)
      (i32.clz
       (local.tee $3
        (select
         (i32.const 1073741820)
         (local.get $2)
         (i32.ge_u
          (local.get $2)
          (i32.const 1073741820)))))))
    (local.set $3
     (i32.xor
      (i32.shr_u
       (local.get $3)
       (i32.sub
        (local.get $2)
        (i32.const 4)))
      (i32.const 16)))
    (i32.sub
     (local.get $2)
     (i32.const 7)))))
 (local.set $2
  (i32.load offset=8
   (local.get $1)))
 (if
  (local.tee $4
   (i32.load offset=4
    (local.get $1)))
  (then
   (i32.store offset=8
    (local.get $4)
    (local.get $2)))
  (else))
 (if
  (local.get $2)
  (then
   (i32.store offset=4
    (local.get $2)
    (local.get $4)))
  (else))
 (if
  (i32.eq
   (local.get $1)
   (i32.load offset=96
    (local.tee $4
     (i32.add
      (local.get $0)
      (i32.shl
       (i32.add
        (i32.shl
         (local.get $5)
         (i32.const 4))
        (local.get $3))
       (i32.const 2))))))
  (then
   (i32.store offset=96
    (local.get $4)
    (local.get $2))
   (if
    (i32.eqz
     (local.get $2))
    (then
     (i32.store offset=4
      (local.tee $1
       (i32.add
        (local.get $0)
        (i32.shl
         (local.get $5)
         (i32.const 2))))
      (local.tee $1
       (i32.and
        (i32.load offset=4
         (local.get $1))
        (i32.rotl
         (i32.const -2)
         (local.get $3)))))
     (if
      (i32.eqz
       (local.get $1))
      (then
       (i32.store
        (local.get $0)
        (i32.and
         (i32.load
          (local.get $0))
         (i32.rotl
          (i32.const -2)
          (local.get $5)))))
      (else)))
    (else)))
  (else)))
(func $tlsf/insertBlock (param $0 i32) (param $1 i32)
 (local $2 i32)
 (local $3 i32)
 (local $4 i32)
 (local $5 i32)
 (local.set $3
  (local.tee $4
   (i32.load
    (local.get $1))))
 (if
  (i32.and
   (local.tee $4
    (i32.load
     (local.tee $2
      (i32.add
       (local.tee $5
        (i32.add
         (local.get $1)
         (i32.const 4)))
       (i32.and
        (local.get $4)
        (i32.const -4))))))
   (i32.const 1))
  (then
   (call $tlsf/removeBlock
    (local.get $0)
    (local.get $2))
   (i32.store
    (local.get $1)
    (local.tee $3
     (i32.add
      (i32.add
       (local.get $3)
       (i32.const 4))
      (i32.and
       (local.get $4)
       (i32.const -4)))))
   (local.set $4
    (i32.load
     (local.tee $2
      (i32.add
       (i32.and
        (i32.load
         (local.get $1))
        (i32.const -4))
       (local.get $5))))))
  (else))
 (if
  (i32.and
   (local.get $3)
   (i32.const 2))
  (then
   (local.set $5
    (i32.load
     (local.tee $1
      (i32.load
       (i32.sub
        (local.get $1)
        (i32.const 4))))))
   (call $tlsf/removeBlock
    (local.get $0)
    (local.get $1))
   (i32.store
    (local.get $1)
    (local.tee $3
     (i32.add
      (i32.add
       (local.get $5)
       (i32.const 4))
      (i32.and
       (local.get $3)
       (i32.const -4))))))
  (else))
 (i32.store
  (local.get $2)
  (i32.or
   (local.get $4)
   (i32.const 2)))
 (i32.store
  (i32.sub
   (local.get $2)
   (i32.const 4))
  (local.get $1))
 (local.set $3
  (if (result i32)
   (i32.lt_u
    (local.tee $2
     (i32.and
      (local.get $3)
      (i32.const -4)))
    (i32.const 256))
   (then
    (local.set $2
     (i32.shr_u
      (local.get $2)
      (i32.const 4)))
    (i32.const 0))
   (else
    (local.set $3
     (i32.sub
      (i32.const 31)
      (i32.clz
       (local.tee $2
        (select
         (i32.const 1073741820)
         (local.get $2)
         (i32.ge_u
          (local.get $2)
          (i32.const 1073741820)))))))
    (local.set $2
     (i32.xor
      (i32.shr_u
       (local.get $2)
       (i32.sub
        (local.get $3)
        (i32.const 4)))
      (i32.const 16)))
    (i32.sub
     (local.get $3)
     (i32.const 7)))))
 (local.set $4
  (i32.load offset=96
   (i32.add
    (local.get $0)
    (i32.shl
     (i32.add
      (i32.shl
       (local.get $3)
       (i32.const 4))
      (local.get $2))
     (i32.const 2)))))
 (i32.store offset=4
  (local.get $1)
  (i32.const 0))
 (i32.store offset=8
  (local.get $1)
  (local.get $4))
 (if
  (local.get $4)
  (then
   (i32.store offset=4
    (local.get $4)
    (local.get $1)))
  (else))
 (i32.store offset=96
  (i32.add
   (local.get $0)
   (i32.shl
    (i32.add
     (i32.shl
      (local.get $3)
      (i32.const 4))
     (local.get $2))
    (i32.const 2)))
  (local.get $1))
 (i32.store
  (local.get $0)
  (i32.or
   (i32.load
    (local.get $0))
   (i32.shl
    (i32.const 1)
    (local.get $3))))
 (i32.store offset=4
  (local.tee $0
   (i32.add
    (local.get $0)
    (i32.shl
     (local.get $3)
     (i32.const 2))))
  (i32.or
   (i32.load offset=4
    (local.get $0))
   (i32.shl
    (i32.const 1)
    (local.get $2)))))
(func $tlsf/addMemory (param $0 i32) (param $1 i32) (param $2 i64)
 (local $3 i32)
 (local $4 i32)
 (local $5 i32)
 (if
  (select
   (local.tee $4
    (i32.load offset=1568
     (local.get $0)))
   (i32.const 0)
   (i32.eq
    (local.tee $3
     (i32.sub
      (local.tee $1
       (i32.sub
        (i32.and
         (i32.add
          (local.get $1)
          (i32.const 19))
         (i32.const -16))
        (i32.const 4)))
      (i32.const 16)))
    (local.get $4)))
  (then
   (local.set $5
    (i32.load
     (local.get $4)))
   (local.set $1
    (local.get $3)))
  (else))
 (if
  (i32.lt_u
   (local.tee $3
    (i32.sub
     (i32.and
      (i32.wrap_i64
       (local.get $2))
      (i32.const -16))
     (local.get $1)))
   (i32.const 20))
  (then
   (return))
  (else))
 (i32.store
  (local.get $1)
  (i32.or
   (i32.and
    (local.get $5)
    (i32.const 2))
   (i32.or
    (local.tee $3
     (i32.sub
      (local.get $3)
      (i32.const 8)))
    (i32.const 1))))
 (i32.store offset=4
  (local.get $1)
  (i32.const 0))
 (i32.store offset=8
  (local.get $1)
  (i32.const 0))
 (i32.store
  (local.tee $3
   (i32.add
    (i32.add
     (local.get $1)
     (i32.const 4))
    (local.get $3)))
  (i32.const 2))
 (i32.store offset=1568
  (local.get $0)
  (local.get $3))
 (call $tlsf/insertBlock
  (local.get $0)
  (local.get $1)))
(func $tlsf/initialize
 (local $0 i32)
 (local $1 i32)
 (local $2 i32)
 (local.set $0
  (i32.and
   (i32.add
    (i32.const 10032)
    (i32.const 15))
   (i32.const -16)))
 (if
  (if (result i32)
   (i32.lt_s
    (local.tee $1
     (memory.size))
    (local.tee $2
     (i32.shr_u
      (i32.and
       (i32.add
        (local.get $0)
        (i32.const 67107))
       (i32.const -65536))
      (i32.const 16))))
   (then
    (i32.lt_s
     (memory.grow
      (i32.sub
       (local.get $2)
       (local.get $1)))
     (i32.const 0)))
   (else
    (i32.const 0)))
  (then
   (unreachable))
  (else))
 (i32.store
  (local.get $0)
  (i32.const 0))
 (i32.store offset=1568
  (local.get $0)
  (i32.const 0))
 (local.set $1
  (i32.const 0))
 (loop $label1
  (if
   (i32.lt_u
    (local.get $1)
    (i32.const 23))
   (then
    (i32.store offset=4
     (i32.add
      (local.get $0)
      (i32.shl
       (local.get $1)
       (i32.const 2)))
     (i32.const 0))
    (local.set $2
     (i32.const 0))
    (loop $label
     (if
      (i32.lt_u
       (local.get $2)
       (i32.const 16))
      (then
       (i32.store offset=96
        (i32.add
         (local.get $0)
         (i32.shl
          (i32.add
           (i32.shl
            (local.get $1)
            (i32.const 4))
           (local.get $2))
          (i32.const 2)))
        (i32.const 0))
       (local.set $2
        (i32.add
         (local.get $2)
         (i32.const 1)))
       (br $label))
      (else)))
    (local.set $1
     (i32.add
      (local.get $1)
      (i32.const 1)))
    (br $label1))
   (else)))
 (call $tlsf/addMemory
  (local.get $0)
  (i32.add
   (local.get $0)
   (i32.const 1572))
  (i64.shl
   (i64.extend_i32_s
    (memory.size))
   (i64.const 16)))
 (global.set $tlsf/ROOT
  (local.get $0)))
(func $moonbit.malloc (param $0 i32) (result i32)
 (local $1 i32)
 (local $2 i32)
 (local $3 i32)
 (local $4 i32)
 (if
  (i32.eqz
   (global.get $tlsf/ROOT))
  (then
   (call $tlsf/initialize))
  (else))
 (if
  (i32.gt_u
   (local.get $0)
   (i32.const 1073741820))
  (then
   (unreachable))
  (else))
 (if
  (i32.eqz
   (local.tee $0
    (call $tlsf/searchBlock
     (local.tee $2
      (global.get $tlsf/ROOT))
     (local.tee $1
      (if (result i32)
       (i32.le_u
        (local.get $0)
        (i32.const 12))
       (then
        (i32.const 12))
       (else
        (i32.sub
         (i32.and
          (i32.add
           (local.get $0)
           (i32.const 19))
          (i32.const -16))
         (i32.const 4))))))))
  (then
   (local.tee $0
    (memory.size))
   (if
    (i32.lt_s
     (memory.grow
      (select
       (if (result i32)
        (i32.ge_u
         (local.get $1)
         (i32.const 256))
        (then
         (if (result i32)
          (i32.lt_u
           (local.get $1)
           (i32.const 536870910))
          (then
           (i32.sub
            (i32.add
             (local.get $1)
             (i32.shl
              (i32.const 1)
              (i32.sub
               (i32.const 27)
               (i32.clz
                (local.get $1)))))
            (i32.const 1)))
          (else
           (local.get $1))))
        (else
         (local.get $1)))
       (local.tee $3
        (i32.shr_u
         (i32.and
          (i32.add
           (i32.add
            (i32.const 4)
            (i32.shl
             (i32.load offset=1568
              (local.get $2))
             (i32.ne
              (i32.sub
               (i32.shl
                (local.get $0)
                (i32.const 16))
               (i32.const 4)))))
           (i32.const 65535))
          (i32.const -65536))
         (i32.const 16)))
       (i32.gt_s
        (local.get $0)
        (local.get $3))))
     (i32.const 0))
    (then
     (if
      (i32.lt_s
       (memory.grow
        (local.get $3))
       (i32.const 0))
      (then
       (unreachable))
      (else)))
    (else))
   (call $tlsf/addMemory
    (local.get $2)
    (i32.shl
     (local.get $0)
     (i32.const 16))
    (i64.shl
     (i64.extend_i32_s
      (memory.size))
     (i64.const 16)))
   (local.set $0
    (call $tlsf/searchBlock
     (local.get $2)
     (local.get $1))))
  (else))
 (call $tlsf/removeBlock
  (local.get $2)
  (local.get $0))
 (if
  (i32.ge_u
   (local.tee $4
    (i32.sub
     (i32.and
      (local.tee $3
       (i32.load
        (local.get $0)))
      (i32.const -4))
     (local.get $1)))
   (i32.const 16))
  (then
   (i32.store
    (local.get $0)
    (i32.or
     (local.get $1)
     (i32.and
      (local.get $3)
      (i32.const 2))))
   (i32.store
    (local.tee $1
     (i32.add
      (i32.add
       (local.get $0)
       (i32.const 4))
      (local.get $1)))
    (i32.or
     (i32.sub
      (local.get $4)
      (i32.const 4))
     (i32.const 1)))
   (call $tlsf/insertBlock
    (local.get $2)
    (local.get $1)))
  (else
   (i32.store
    (local.get $0)
    (i32.and
     (local.get $3)
     (i32.const -2)))
   (i32.store
    (i32.add
     (local.tee $1
      (i32.add
       (local.get $0)
       (i32.const 4)))
     (local.tee $2
      (i32.and
       (i32.load
        (local.get $0))
       (i32.const -4))))
    (i32.and
     (i32.load
      (i32.add
       (local.get $1)
       (local.get $2)))
     (i32.const -3)))))
 (i32.add
  (local.get $0)
  (i32.const 4)))
(func $moonbit.free (param $0 i32)
 (local $1 i32)
 (local $2 i32)
 (if
  (i32.gt_u
   (i32.const 10032)
   (local.get $0))
  (then
   (return))
  (else))
 (if
  (i32.eqz
   (global.get $tlsf/ROOT))
  (then
   (call $tlsf/initialize))
  (else))
 (local.set $2
  (global.get $tlsf/ROOT))
 (local.set $1
  (i32.sub
   (local.get $0)
   (i32.const 4)))
 (if
  (if (result i32)
   (select
    (i32.and
     (local.get $0)
     (i32.const 15))
    (i32.const 1)
    (local.get $0))
   (then
    (i32.const 1))
   (else
    (i32.and
     (i32.load
      (local.get $1))
     (i32.const 1))))
  (then
   (unreachable))
  (else))
 (i32.store
  (local.get $1)
  (i32.or
   (i32.load
    (local.get $1))
   (i32.const 1)))
 (call $tlsf/insertBlock
  (local.get $2)
  (local.get $1)))
(func $moonbit.gc.malloc (param $n i32) (result i32)
 (local $result i32)
 (i32.store
  (local.tee $result
   (call $moonbit.malloc
    (i32.add
     (i32.const 8)
     (local.get $n))))
  (i32.const 1))
 (local.get $result))
(func $moonbit.array_length (param $arr i32) (result i32)
 (i32.and
  (i32.load offset=4
   (local.get $arr))
  (i32.const 268435455)))
(func $moonbit.incref (param $ptr i32)
 (local $count i32)
 (if
  (i32.ge_s
   (local.tee $count
    (i32.load
     (local.get $ptr)))
   (i32.const 0))
  (then
   (i32.store
    (local.get $ptr)
    (i32.add
     (local.get $count)
     (i32.const 1))))
  (else)))
(func $moonbit.decref (param $ptr i32)
 (local $count i32)
 (if
  (i32.gt_s
   (local.tee $count
    (i32.load
     (local.get $ptr)))
   (i32.const 1))
  (then
   (i32.store
    (local.get $ptr)
    (i32.sub
     (local.get $count)
     (i32.const 1))))
  (else
   (if
    (i32.eq
     (local.get $count)
     (i32.const 1))
    (then
     (call $moonbit.gc.free
      (local.get $ptr)))
    (else)))))
(func $moonbit.gc.free (param $ptr i32)
 (local $parent i32)
 (local $curr_child_offset i32)
 (local $remaining_children_count i32)
 (local $n_ptr_fields i32)
 (local $ptr_fields_offset i32)
 (local $kind i32)
 (local $ref_array_kind i32)
 (local $vt_ptr i32)
 (local $vt_ptr_index i32)
 (local $vt_ptr_fields_offset i32)
 (local $vt_n_ptr_fields i32)
 (local $vt_header i32)
 (local $meta i32)
 (local $next i32)
 (local $addr_of_next i32)
 (local $len i32)
 (local $count i32)
 (loop $handle_new_object
  (local.set $kind
   (i32.shr_u
    (local.tee $meta
     (i32.load offset=4
      (local.get $ptr)))
    (i32.const 30)))
  (block $cond_has_children
   (if
    (i32.eq
     (i32.const 0)
     (local.get $kind))
    (then
     (if
      (i32.eqz
       (local.tee $n_ptr_fields
        (i32.and
         (i32.shr_u
          (local.get $meta)
          (i32.const 8))
         (i32.const 2047))))
      (then)
      (else
       (local.set $ptr_fields_offset
        (i32.and
         (i32.shr_u
          (local.get $meta)
          (i32.const 19))
         (i32.const 2047)))
       (local.set $curr_child_offset
        (local.get $ptr_fields_offset))
       (local.set $remaining_children_count
        (local.get $n_ptr_fields))
       (br $cond_has_children))))
    (else
     (if
      (i32.eq
       (i32.const 2)
       (local.get $kind))
      (then
       (local.set $ref_array_kind
        (i32.and
         (i32.const 3)
         (i32.shr_u
          (local.get $meta)
          (i32.const 28))))
       (if
        (i32.eq
         (i32.const 1)
         (local.get $ref_array_kind))
        (then
         (local.set $len
          (i32.and
           (local.get $meta)
           (i32.const 268435455)))
         (local.set $vt_ptr
          (i32.add
           (local.get $ptr)
           (i32.const 8)))
         (local.set $vt_header
          (i32.load
           (local.get $vt_ptr)))
         (local.set $vt_n_ptr_fields
          (i32.and
           (i32.const 2047)
           (i32.shr_u
            (local.get $vt_header)
            (i32.const 8))))
         (local.set $vt_ptr_fields_offset
          (i32.and
           (i32.const 2047)
           (i32.shr_u
            (local.get $vt_header)
            (i32.const 19))))
         (local.set $vt_ptr
          (i32.add
           (local.get $vt_ptr)
           (i32.const 4)))
         (loop $vt_elems_loop
          (if
           (i32.gt_s
            (local.get $len)
            (i32.const 0))
           (then
            (local.set $len
             (i32.sub
              (local.get $len)
              (i32.const 1)))
            (local.set $vt_ptr_index
             (i32.const 0))
            (local.set $vt_ptr
             (i32.add
              (local.get $vt_ptr)
              (i32.mul
               (local.get $vt_ptr_fields_offset)
               (i32.const 4))))
            (loop $vt_ptrs_loop
             (if
              (i32.lt_s
               (local.get $vt_ptr_index)
               (local.get $vt_n_ptr_fields))
              (then
               (i32.load
                (local.get $vt_ptr))
               (if
                (i32.ne
                 (i32.const 0))
                (then
                 (call $moonbit.decref
                  (i32.load
                   (local.get $vt_ptr))))
                (else))
               (local.set $vt_ptr_index
                (i32.add
                 (local.get $vt_ptr_index)
                 (i32.const 1)))
               (local.set $vt_ptr
                (i32.add
                 (local.get $vt_ptr)
                 (i32.const 4)))
               (br $vt_ptrs_loop))
              (else)))
            (br $vt_elems_loop))
           (else))))
        (else
         (local.set $len
          (i32.and
           (local.get $meta)
           (i32.const 268435455)))
         (if
          (i32.gt_s
           (local.get $len)
           (i32.const 0))
          (then
           (local.set $curr_child_offset
            (i32.shr_u
             (i32.const 8)
             (i32.const 2)))
           (local.set $remaining_children_count
            (local.get $len))
           (br $cond_has_children))
          (else)))))
      (else
       (if
        (i32.eq
         (i32.const 1)
         (local.get $kind))
        (then)
        (else
         (unreachable)))))))
   (call $moonbit.free
    (local.get $ptr))
   (if
    (i32.eqz
     (local.get $parent))
    (then
     (return))
    (else))
   (local.set $curr_child_offset
    (i32.load
     (local.get $parent)))
   (local.set $remaining_children_count
    (i32.load offset=4
     (local.get $parent)))
   (local.set $ptr
    (local.get $parent))
   (local.set $parent
    (i32.load
     (i32.add
      (local.get $ptr)
      (i32.mul
       (local.get $curr_child_offset)
       (i32.const 4)))))
   (local.set $curr_child_offset
    (i32.add
     (local.get $curr_child_offset)
     (i32.const 1))))
  (loop $process_children
   (loop $process_children_loop
    (if
     (i32.gt_s
      (local.get $remaining_children_count)
      (i32.const 0))
     (then
      (local.set $remaining_children_count
       (i32.sub
        (local.get $remaining_children_count)
        (i32.const 1)))
      (if
       (i32.eqz
        (local.tee $next
         (i32.load
          (local.tee $addr_of_next
           (i32.add
            (local.get $ptr)
            (i32.mul
             (local.get $curr_child_offset)
             (i32.const 4)))))))
       (then
        (local.set $curr_child_offset
         (i32.add
          (local.get $curr_child_offset)
          (i32.const 1)))
        (br $process_children_loop))
       (else))
      (if
       (i32.gt_s
        (local.tee $count
         (i32.load
          (local.get $next)))
        (i32.const 1))
       (then
        (i32.store
         (local.get $next)
         (i32.sub
          (local.get $count)
          (i32.const 1))))
       (else
        (if
         (i32.eq
          (local.get $count)
          (i32.const 1))
         (then
          (if
           (i32.eq
            (local.get $remaining_children_count)
            (i32.const 0))
           (then
            (call $moonbit.free
             (local.get $ptr)))
           (else
            (i32.store
             (local.get $ptr)
             (local.get $curr_child_offset))
            (i32.store offset=4
             (local.get $ptr)
             (local.get $remaining_children_count))
            (i32.store
             (local.get $addr_of_next)
             (local.get $parent))
            (local.set $parent
             (local.get $ptr))))
          (local.set $ptr
           (local.get $next))
          (br $handle_new_object))
         (else))))
      (local.set $curr_child_offset
       (i32.add
        (local.get $curr_child_offset)
        (i32.const 1)))
      (br $process_children_loop))
     (else)))
   (call $moonbit.free
    (local.get $ptr))
   (if
    (i32.eqz
     (local.get $parent))
    (then
     (return))
    (else))
   (local.set $curr_child_offset
    (i32.load
     (local.get $parent)))
   (local.set $remaining_children_count
    (i32.load offset=4
     (local.get $parent)))
   (local.set $ptr
    (local.get $parent))
   (local.set $parent
    (i32.load
     (i32.add
      (local.get $ptr)
      (i32.mul
       (local.get $curr_child_offset)
       (i32.const 4)))))
   (local.set $curr_child_offset
    (i32.add
     (local.get $curr_child_offset)
     (i32.const 1)))
   (br $process_children))
  (unreachable)))
(table $moonbit.global 0 0 funcref )
(elem
 (table $moonbit.global) (offset (i32.const 0))
 funcref
 )
(global $applejag/firefly-jam-2026.racebattle_world
 (mut i32)
 (i32.const 0)
)
(func $$applejag/firefly-jam-2026.render_wrapper/1
 (drop
  (call $applejag/firefly-jam-2026.render)))
(export "render" (func $$applejag/firefly-jam-2026.render_wrapper/1))
(func $applejag/firefly-jam-2026.render (result i32)
 (i32.const 0))
(func $$applejag/firefly-jam-2026.update_wrapper/2
 (drop
  (call $applejag/firefly-jam-2026.update)))
(export "update" (func $$applejag/firefly-jam-2026.update_wrapper/2))
(func $applejag/firefly-jam-2026.update (result i32)
 (local $val/76 i32)
 (local $*field/77 i32)
 (call $moonbit.incref
  (local.tee $val/76
   (local.tee $*field/77
    (i32.load offset=8
     (global.get $applejag/firefly-jam-2026.racebattle_world)))))
 (call $@applejag/firefly-jam-2026/racebattle.World::update
  (local.get $val/76)))
(func $$applejag/firefly-jam-2026.boot_wrapper/3
 (drop
  (call $applejag/firefly-jam-2026.boot)))
(export "boot" (func $$applejag/firefly-jam-2026.boot_wrapper/3))
(func $applejag/firefly-jam-2026.boot (result i32)
 (local $*bind/23 i32)
 (local $*tmp/73 i32)
 (local $player_system/74 i32)
 (local $camera/75 i32)
 (local $*old/78 i32)
 (local $*field/79 i32)
 (local $*field/80 i32)
 (local $*field/81 i32)
 (local $ptr/98 i32)
 (local.set $player_system/74
  (local.tee $*field/80
   (i32.load offset=8
    (local.tee $*bind/23
     (local.tee $*field/81
      (i32.load offset=8
       (global.get $applejag/firefly-jam-2026.racebattle_world)))))))
 (local.set $camera/75
  (local.tee $*field/79
   (i32.load offset=12
    (local.get $*bind/23))))
 (call $moonbit.incref
  (local.get $player_system/74))
 (call $moonbit.incref
  (local.get $camera/75))
 (i32.store offset=4
  (local.tee $ptr/98
   (call $moonbit.gc.malloc
    (i32.const 8)))
  (i32.const 1049088))
 (i32.store offset=12
  (local.get $ptr/98)
  (local.get $camera/75))
 (i32.store offset=8
  (local.get $ptr/98)
  (local.get $player_system/74))
 (local.set $*tmp/73
  (local.get $ptr/98))
 (call $moonbit.decref
  (local.tee $*old/78
   (i32.load offset=8
    (global.get $applejag/firefly-jam-2026.racebattle_world))))
 (i32.store offset=8
  (global.get $applejag/firefly-jam-2026.racebattle_world)
  (local.get $*tmp/73))
 (i32.const 0))
(func $@applejag/firefly-jam-2026/racebattle.World::update (param $self/22 i32) (result i32)
 (local $camera/72 i32)
 (local $*field/82 i32)
 (call $moonbit.incref
  (local.tee $camera/72
   (local.tee $*field/82
    (i32.load offset=12
     (local.get $self/22)))))
 (call $@applejag/firefly-jam-2026/racebattle.Camera::update
  (local.get $camera/72)
  (local.get $self/22)))
(func $@applejag/firefly-jam-2026/racebattle.Camera::update (param $self/21 i32) (param $world/20 i32) (result i32)
 (local $my_peer/0/18 i32)
 (local $maybe_idx/19 i64)
 (local $*tmp/0/65 f32)
 (local $*tmp/1/65 f32)
 (local $pos/0/66 f32)
 (local $pos/1/66 f32)
 (local $*tmp/0/67 f32)
 (local $*tmp/1/67 f32)
 (local $players_pos/68 i32)
 (local $player_system/69 i32)
 (local $players_peer/70 i32)
 (local $player_system/71 i32)
 (local $*field/83 i32)
 (local $*field/84 i32)
 (local $*field/85 i32)
 (local $*field/86 i32)
 (local $*cnt/91 i32)
 (local $*field/92 i32)
 (local $*new_cnt/93 i32)
 (local $*cnt/94 i32)
 (local $*field/95 i32)
 (local $*field/96 i32)
 (local $*new_cnt/97 i32)
 (local.set $my_peer/0/18
  (call $firefly/firefly.get_me))
 (call $moonbit.incref
  (local.tee $players_peer/70
   (local.tee $*field/85
    (i32.load offset=8
     (local.tee $player_system/71
      (local.tee $*field/86
       (i32.load offset=8
        (local.get $world/20))))))))
 (if (result i32)
  (i64.eq
   (local.tee $maybe_idx/19
    (call $ReadOnlyArray::search|@firefly/firefly.Peer|
     (local.get $players_peer/70)
     (local.get $my_peer/0/18)))
   (i64.const 4294967296))
  (then
   (call $moonbit.decref
    (local.get $self/21))
   (call $moonbit.decref
    (local.get $world/20))
   (i32.const 0))
  (else
   (f32.load offset=8
    (local.get $self/21))
   (local.set $pos/1/66
    (f32.load offset=12
     (local.get $self/21)))
   (local.set $pos/0/66)
   (call $@applejag/firefly-jam-2026/util.Vec2::new
    (f32.const 0x1.ep+6)
    (f32.const 0x1.4p+6))
   (local.set $*tmp/1/67)
   (local.set $*tmp/0/67)
   (call $@moonbitlang/core/builtin.Add::@applejag/firefly-jam-2026/util.Vec2::add
    (local.get $pos/0/66)
    (local.get $pos/1/66)
    (local.get $*tmp/0/67)
    (local.get $*tmp/1/67))
   (local.set $*tmp/1/65)
   (local.set $*tmp/0/65)
   (f32.store offset=8
    (local.get $self/21)
    (local.get $*tmp/0/65))
   (f32.store offset=12
    (local.get $self/21)
    (local.get $*tmp/1/65))
   (call $moonbit.decref
    (local.get $self/21))
   (local.set $*field/84
    (i32.load offset=8
     (local.get $world/20)))
   (if
    (i32.gt_s
     (local.tee $*cnt/91
      (i32.load
       (local.get $world/20)))
     (i32.const 1))
    (then
     (local.set $*new_cnt/93
      (i32.sub
       (local.get $*cnt/91)
       (i32.const 1)))
     (i32.store
      (local.get $world/20)
      (local.get $*new_cnt/93))
     (call $moonbit.incref
      (local.get $*field/84)))
    (else
     (if
      (i32.eq
       (local.get $*cnt/91)
       (i32.const 1))
      (then
       (call $moonbit.decref
        (local.tee $*field/92
         (i32.load offset=12
          (local.get $world/20))))
       (call $moonbit.free
        (local.get $world/20)))
      (else))))
   (local.set $*field/83
    (i32.load offset=12
     (local.tee $player_system/69
      (local.get $*field/84))))
   (if
    (i32.gt_s
     (local.tee $*cnt/94
      (i32.load
       (local.get $player_system/69)))
     (i32.const 1))
    (then
     (local.set $*new_cnt/97
      (i32.sub
       (local.get $*cnt/94)
       (i32.const 1)))
     (i32.store
      (local.get $player_system/69)
      (local.get $*new_cnt/97))
     (call $moonbit.incref
      (local.get $*field/83)))
    (else
     (if
      (i32.eq
       (local.get $*cnt/94)
       (i32.const 1))
      (then
       (call $moonbit.decref
        (local.tee $*field/96
         (i32.load offset=16
          (local.get $player_system/69))))
       (call $moonbit.decref
        (local.tee $*field/95
         (i32.load offset=8
          (local.get $player_system/69))))
       (call $moonbit.free
        (local.get $player_system/69)))
      (else))))
   (call $moonbit.decref
    (local.tee $players_pos/68
     (local.get $*field/83)))
   (i32.const 0))))
(func $@moonbitlang/core/builtin.Default::@applejag/firefly-jam-2026/racebattle.PlayerSystem::default (result i32)
 (local $*tmp/62 i32)
 (local $*tmp/63 i32)
 (local $*tmp/64 i32)
 (local $ptr/99 i32)
 (local.set $*tmp/62
  (call $@moonbitlang/core/builtin.Default::ReadOnlyArray::default|@firefly/firefly.Peer|))
 (local.set $*tmp/63
  (call $@moonbitlang/core/builtin.Default::FixedArray::default|@applejag/firefly-jam-2026/util.Vec2|))
 (local.set $*tmp/64
  (call $@moonbitlang/core/builtin.Default::FixedArray::default|@firefly/firefly.Angle|))
 (i32.store offset=4
  (local.tee $ptr/99
   (call $moonbit.gc.malloc
    (i32.const 12)))
  (i32.const 1049344))
 (i32.store offset=16
  (local.get $ptr/99)
  (local.get $*tmp/64))
 (i32.store offset=12
  (local.get $ptr/99)
  (local.get $*tmp/63))
 (i32.store offset=8
  (local.get $ptr/99)
  (local.get $*tmp/62))
 (local.get $ptr/99))
(func $@moonbitlang/core/builtin.Default::@applejag/firefly-jam-2026/racebattle.World::default (result i32)
 (local $*tmp/60 i32)
 (local $*tmp/61 i32)
 (local $ptr/100 i32)
 (local.set $*tmp/60
  (call $@moonbitlang/core/builtin.Default::@applejag/firefly-jam-2026/racebattle.PlayerSystem::default))
 (local.set $*tmp/61
  (call $@moonbitlang/core/builtin.Default::@applejag/firefly-jam-2026/racebattle.Camera::default))
 (i32.store offset=4
  (local.tee $ptr/100
   (call $moonbit.gc.malloc
    (i32.const 8)))
  (i32.const 1049088))
 (i32.store offset=12
  (local.get $ptr/100)
  (local.get $*tmp/61))
 (i32.store offset=8
  (local.get $ptr/100)
  (local.get $*tmp/60))
 (local.get $ptr/100))
(func $@moonbitlang/core/builtin.Default::@applejag/firefly-jam-2026/racebattle.Camera::default (result i32)
 (local $*tmp/0/59 f32)
 (local $*tmp/1/59 f32)
 (local $ptr/101 i32)
 (call $@moonbitlang/core/builtin.Default::@applejag/firefly-jam-2026/util.Vec2::default)
 (local.set $*tmp/1/59)
 (local.set $*tmp/0/59)
 (i32.store offset=4
  (local.tee $ptr/101
   (call $moonbit.gc.malloc
    (i32.const 8)))
  (i32.const 2097152))
 (f32.store offset=8
  (local.get $ptr/101)
  (local.get $*tmp/0/59))
 (f32.store offset=12
  (local.get $ptr/101)
  (local.get $*tmp/1/59))
 (local.get $ptr/101))
(func $@moonbitlang/core/builtin.Add::@applejag/firefly-jam-2026/util.Vec2::add (param $self/0/16 f32) (param $self/1/16 f32) (param $other/0/17 f32) (param $other/1/17 f32) (result f32) (result f32)
 (local $*tmp/53 f32)
 (local $*tmp/54 f32)
 (local $y/55 f32)
 (local $y/56 f32)
 (local $x/57 f32)
 (local $x/58 f32)
 (local.set $x/57
  (local.get $self/0/16))
 (local.set $x/58
  (local.get $other/0/17))
 (local.set $*tmp/53
  (f32.add
   (local.get $x/57)
   (local.get $x/58)))
 (local.set $y/55
  (local.get $self/1/16))
 (local.set $y/56
  (local.get $other/1/17))
 (local.set $*tmp/54
  (f32.add
   (local.get $y/55)
   (local.get $y/56)))
 (local.get $*tmp/53)
 (local.get $*tmp/54))
(func $@applejag/firefly-jam-2026/util.Vec2::new (param $x/14 f32) (param $y/15 f32) (result f32) (result f32)
 (local.get $x/14)
 (local.get $y/15))
(func $@moonbitlang/core/builtin.Default::@applejag/firefly-jam-2026/util.Vec2::default (result f32) (result f32)
 (f32.const 0x0p+0)
 (f32.const 0x0p+0))
(func $firefly/firefly.get_me (result i32)
 (local $*tmp/52 i32)
 (local.tee $*tmp/52
  (call $firefly/firefly/internal/ffi.get_me)))
(func $Ref::new|@applejag/firefly-jam-2026/racebattle.World| (param $x/11 i32) (result i32)
 (local $ptr/102 i32)
 (i32.store offset=4
  (local.tee $ptr/102
   (call $moonbit.gc.malloc
    (i32.const 4)))
  (i32.const 1048832))
 (i32.store offset=8
  (local.get $ptr/102)
  (local.get $x/11))
 (local.get $ptr/102))
(func $@moonbitlang/core/builtin.Default::ReadOnlyArray::default|@firefly/firefly.Peer| (result i32)
 (local $*tmp/49 i32)
 (local.tee $*tmp/49
  (call $@moonbitlang/core/builtin.Default::FixedArray::default|@firefly/firefly.Peer|)))
(func $@moonbitlang/core/builtin.Default::FixedArray::default|@firefly/firefly.Peer| (result i32)
 (i32.const 10000))
(func $@moonbitlang/core/builtin.Default::FixedArray::default|@firefly/firefly.Angle| (result i32)
 (i32.const 10016))
(func $@moonbitlang/core/builtin.Default::FixedArray::default|@applejag/firefly-jam-2026/util.Vec2| (result i32)
 (i32.const 10000))
(func $ReadOnlyArray::search|@firefly/firefly.Peer| (param $self/9 i32) (param $value/0/10 i32) (result i64)
 (local $*tmp/48 i32)
 (call $FixedArray::search|@firefly/firefly.Peer|
  (local.tee $*tmp/48
   (local.get $self/9))
  (local.get $value/0/10)))
(func $FixedArray::search|@firefly/firefly.Peer| (param $self/7 i32) (param $value/0/8 i32) (result i64)
 (local $*tmp/0/45 i32)
 (local $*tmp/1/45 i32)
 (local $*tmp/2/45 i32)
 (local $*tmp/46 i32)
 (local $*tmp/47 i32)
 (local $*tmp/87 i32)
 (call $moonbit.incref
  (local.get $self/7))
 (local.set $*tmp/46
  (local.get $self/7))
 (local.set $*tmp/87
  (call $moonbit.array_length
   (local.get $self/7)))
 (call $moonbit.decref
  (local.get $self/7))
 (local.set $*tmp/47
  (local.get $*tmp/87))
 (local.get $*tmp/46)
 (i32.const 0)
 (local.set $*tmp/2/45
  (local.get $*tmp/47))
 (local.set $*tmp/1/45)
 (local.tee $*tmp/0/45)
 (local.get $*tmp/1/45)
 (local.get $*tmp/2/45)
 (local.get $value/0/8)
 (call $@moonbitlang/core/builtin.ArrayView::search|@firefly/firefly.Peer|))
(func $@moonbitlang/core/builtin.ArrayView::search|@firefly/firefly.Peer| (param $self/0/3 i32) (param $self/1/3 i32) (param $self/2/3 i32) (param $value/0/5 i32) (result i64)
 (local $*end989/2 i32)
 (local $i/4 i32)
 (local $*p/0/26 i32)
 (local $raw/36 i32)
 (local $raw/37 i32)
 (local $buf/38 i32)
 (local $*tmp/39 i32)
 (local $start/40 i32)
 (local $*tmp/41 i64)
 (local $*tmp/42 i32)
 (local $end/43 i32)
 (local $start/44 i32)
 (local $*tmp/0/88 i32)
 (local $*field/89 i32)
 (local $*arr/103 i32)
 (local $*idx/104 i32)
 (local $*ofs/105 i32)
 (local.set $end/43
  (local.get $self/2/3))
 (local.set $start/44
  (local.get $self/1/3))
 (local.set $*end989/2
  (i32.sub
   (local.get $end/43)
   (local.get $start/44)))
 (block $break:6 (result i64)
  (i32.const 0)
  (loop $loop:6 (param i32) (result i64)
   (local.tee $i/4)
   (local.get $*end989/2)
   (i32.lt_s)
   (if (result i64)
    (then
     (local.set $buf/38
      (local.tee $*field/89
       (local.get $self/0/3)))
     (local.set $*tmp/39
      (i32.add
       (local.tee $start/40
        (local.get $self/1/3))
       (local.get $i/4)))
     (local.get $buf/38)
     (local.set $*idx/104
      (local.get $*tmp/39))
     (local.set $*arr/103)
     (local.set $*ofs/105
      (i32.mul
       (local.get $*idx/104)
       (i32.const 4)))
     (local.set $raw/36
      (local.tee $*p/0/26
       (local.tee $*tmp/0/88
        (i32.load offset=8
         (i32.add
          (i32.add
           (local.get $*arr/103)
           (local.get $*ofs/105))
          (i32.const 0))))))
     (local.set $raw/37
      (local.get $value/0/5))
     (if
      (i32.eq
       (local.get $raw/36)
       (local.get $raw/37))
      (then
       (call $moonbit.decref
        (local.get $self/0/3))
       (local.tee $*tmp/41
        (i64.extend_i32_s
         (local.get $i/4)))
       (br $break:6))
      (else))
     (local.tee $*tmp/42
      (i32.add
       (local.get $i/4)
       (i32.const 1)))
     (br $loop:6))
    (else
     (call $moonbit.decref
      (local.get $self/0/3))
     (i64.const 4294967296))))))
(start $*init*/5)
(func $*init*/5
 (local $*tmp/33 i32)
 (global.set $applejag/firefly-jam-2026.racebattle_world
  (call $Ref::new|@applejag/firefly-jam-2026/racebattle.World|
   (local.tee $*tmp/33
    (call $@moonbitlang/core/builtin.Default::@applejag/firefly-jam-2026/racebattle.World::default)))))
(func $*main*/4)
(export "_start" (func $*main*/4))
